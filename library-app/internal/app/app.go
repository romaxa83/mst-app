package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/romaxa83/mst-app/library-app/internal/config"
	delivery "github.com/romaxa83/mst-app/library-app/internal/delivery/http"
	//kafkaConsumer "github.com/romaxa83/mst-app/library-app/internal/delivery/kafka"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/internal/server"
	"github.com/romaxa83/mst-app/library-app/internal/services"
	"github.com/romaxa83/mst-app/library-app/internal/utils"
	"github.com/romaxa83/mst-app/library-app/pkg/cache"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"github.com/romaxa83/mst-app/library-app/pkg/storage"
	"golang.org/x/text/language"
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

	//ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	//defer cancel()

	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found")
		return
	}
	logger.Info("Load .env file")

	cfg, err := config.Init("configs")

	//logger.Warnf("CONTEXT %+v", ctx)

	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Init config")

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

	// migrate
	if err := models.InitModels(db); err != nil {
		logger.Error("failed init db: %s", err.Error())
	}

	// todo зарефакторить логгер
	// kafka использует zapLogger , а здесь используется logrus
	//zapLoggerConfig := zapLogger.NewLoggerConfig("debug", false, "json")
	//zapL := zapLogger.NewAppLogger(zapLoggerConfig)
	//kafkaProducer := kafkaClient.NewProducer(zapL, cfg.Kafka.Brokers)
	//defer kafkaProducer.Close() // nolint: errcheck
	//logger.Info("Init kafkaProducer")
	//
	//var kafkaConn *kafka.Conn
	//controller, err := kafkaConn.Controller()
	//if err != nil {
	//	logger.Error("kafkaConn.Controller:", err)
	//	return
	//}
	//
	//controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	//logger.Warnf("controllerURI %+v", controllerURI)

	//if cfg.Kafka.InitTopics {
	//	initKafkaTopics(cfg)
	//}

	//k := kafkaMessages.

	//_ = kafkaConsumer.NewAuthorMessageProcessor(cfg)
	//authorMessageProcessor := kafkaConsumer.NewAuthorMessageProcessor(cfg)
	//cg := kafkaClient.NewConsumerGroup(cfg.Kafka.Brokers, cfg.Kafka.GroupID, zapL)
	//go cg.ConsumeTopic(ctx, getConsumerGroupTopics(cfg), kafkaConsumer.PoolSize, authorMessageProcessor.ProcessMessages)

	// cache
	memCache := cache.NewMemoryCache()
	logger.Info("Init memory cache")

	// seeder
	// todo нужно оптимизировать
	//for _, seed := range seed.All() {
	//	if err := seed.Run(db); err != nil {
	//		logger.Errorf("Running seed '%s', failed with error: %s", seed.Name, err)
	//	}
	//}

	storageProvider, err := newStorageProvider(cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Init storage provider [minio]")

	lang := language.English
	if cfg.Locale.Default == language.Russian.String() {
		lang = language.Russian
	}

	locale := utils.NewLocale(i18n.NewBundle(lang))
	logger.Infof(fmt.Sprintf("Init locale, default - [%s]", lang))

	// Services, Repos & API Handlers
	repos := repositories.NewRepositories(db)
	services := services.NewServices(services.Deps{
		Repos:           repos,
		StorageProvider: storageProvider,
		Environment:     cfg.Environment,
		Cache:           memCache,
		CacheTTL:        int64(cfg.CacheTTL.Seconds()),
	})
	handlers := delivery.NewHandler(services, locale)

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

	//if err := db.; err != nil {
	//	logger.Errorf("error occured on db connection close: %v", err)
	//}
}

func newStorageProvider(cfg *config.Config) (storage.Provider, error) {

	client, err := minio.New(cfg.FileStorage.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.FileStorage.AccessKey, cfg.FileStorage.SecretKey, ""),
		//Secure: true,
	})
	if err != nil {
		return nil, err
	}

	provider := storage.NewFileStorage(client, cfg.FileStorage.Bucket, cfg.FileStorage.Endpoint)

	return provider, nil
}
