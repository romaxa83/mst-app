package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/romaxa83/mst-app/library-app/internal/config"
	delivery "github.com/romaxa83/mst-app/library-app/internal/delivery/http"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/internal/server"
	"github.com/romaxa83/mst-app/library-app/internal/services"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Library App API
// @version 1.0
// @description API Server for Library project

// @host localhost:8060
// @BasePath /api

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

	//logger.Warnf("%+v", cfg)

	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Init config")

	//db, err := db.NewClient(
	//	cfg.Postgres.Host,
	//	cfg.Postgres.Port,
	//	cfg.Postgres.Username,
	//	cfg.Postgres.Password,
	//	cfg.Postgres.DBName,
	//	cfg.Postgres.SSLMode)
	//if err != nil {
	//	logger.Error("failed init db: %s", err.Error())
	//}

	// Services, Repos & API Handlers
	repos := repositories.NewRepositories()
	services := services.NewServices(services.Deps{
		Repos: repos})
	handlers := delivery.NewHandler(services)

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Infof("Server started [%s:%s]", cfg.HTTP.Host, cfg.HTTP.Port)
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

	//if err := db.Close(); err != nil {
	//	logger.Errorf("error occured on db connection close: %v", err)
	//}
}
