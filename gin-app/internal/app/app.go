package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/romaxa83/mst-app/gin-app/internal/config"
	delivery "github.com/romaxa83/mst-app/gin-app/internal/delivery/http"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories"
	"github.com/romaxa83/mst-app/gin-app/internal/server"
	"github.com/romaxa83/mst-app/gin-app/internal/services"
	"github.com/romaxa83/mst-app/gin-app/pkg/auth"
	"github.com/romaxa83/mst-app/gin-app/pkg/db"
	"github.com/romaxa83/mst-app/gin-app/pkg/email/smtp"
	"github.com/romaxa83/mst-app/gin-app/pkg/hash"
	"github.com/romaxa83/mst-app/gin-app/pkg/logger"
	"github.com/romaxa83/mst-app/gin-app/pkg/otp"
	"github.com/romaxa83/mst-app/gin-app/pkg/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Gin-App API
// @version 1.0
// @description REST API for Gin-App

// @host localhost:8080
// @BasePath /api/v1/

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

func Run() {
	//logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found")
		return
	}
	logger.Info("Load .env file")

	cfg, err := config.Init("configs")
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Init config")

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	logger.Info("Init hasher")

	db, err := db.NewClient(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode)
	if err != nil {
		logger.Error("failed init db: %s", err.Error())
	}

	emailSender, err := smtp.NewSMTPSender(cfg.SMTP.From, cfg.SMTP.Pass, cfg.SMTP.Host, cfg.SMTP.Port)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Init email sender")

	storageProvider, err := newStorageProvider(cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("Init storageProvider [%s]", cfg.FileStorage.Driver)

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}

	// Init one-time password
	otpGenerator := otp.NewGOTPGenerator()

	logger.Info("Init tokenManager")

	// Services, Repos & API Handlers
	repos := repositories.NewRepositories(db)
	services := services.NewServices(services.Dependencies{
		Repos:                  repos,
		Hasher:                 hasher,
		TokenManager:           tokenManager,
		AccessTokenTTL:         cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL:        cfg.Auth.JWT.RefreshTokenTTL,
		EmailSender:            emailSender,
		EmailConfig:            cfg.Email,
		OtpGenerator:           otpGenerator,
		VerificationCodeLength: cfg.Auth.VerificationCodeLength,
		StorageProvider:        storageProvider,
		Environment:            cfg.Environment,
	})
	handlers := delivery.NewHandler(services, tokenManager)

	services.Files.InitStorageUploaderWorkers(context.Background())

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")
	logger.Info(fmt.Sprintf("ENV [%s]", cfg.Environment))

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := db.Close(); err != nil {
		logger.Errorf("error occured on db connection close: %v", err)
	}
}

func newStorageProvider(cfg *config.Config) (storage.Provider, error) {

	provider := storage.NewFileStorage(
		cfg.FtpStorage.BaseUrl,
		cfg.FtpStorage.Host,
		cfg.FtpStorage.Port,
		cfg.FtpStorage.Username,
		cfg.FtpStorage.Password,
		cfg.FtpStorage.TimeOut,
	)

	return provider, nil
}
