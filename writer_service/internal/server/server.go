package server

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/romaxa83/mst-app/pkg/interceptors"
	kafkaClient "github.com/romaxa83/mst-app/pkg/kafka"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/pkg/postgres"
	"github.com/romaxa83/mst-app/pkg/tracing"
	"github.com/romaxa83/mst-app/writer_service/config"
	"github.com/romaxa83/mst-app/writer_service/internal/metrics"
	kafkaConsumer "github.com/romaxa83/mst-app/writer_service/internal/product/delivery/kafka"
	"github.com/romaxa83/mst-app/writer_service/internal/product/repository"
	"github.com/romaxa83/mst-app/writer_service/internal/product/service"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log       logger.Logger
	cfg       *config.Config
	v         *validator.Validate
	kafkaConn *kafka.Conn
	ps        *service.ProductService
	im        interceptors.InterceptorManager
	pgConn    *pgxpool.Pool
	metrics   *metrics.WriterServiceMetrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)
	s.metrics = metrics.NewWriterServiceMetrics(s.cfg)

	pgxConn, err := postgres.NewPgxConn(s.cfg.Postgresql)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.pgConn = pgxConn
	s.log.Infof("postgres connected: %v", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()

	kafkaProducer := kafkaClient.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errcheck

	productRepo := repository.NewProductRepository(s.log, s.cfg, pgxConn)
	s.ps = service.NewProductService(s.log, s.cfg, productRepo, kafkaProducer)
	productMessageProcessor := kafkaConsumer.NewProductMessageProcessor(s.log, s.cfg, s.v, s.ps, s.metrics)

	s.log.Info("Starting Writer Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID, s.log)
	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), kafkaConsumer.PoolSize, productMessageProcessor.ProcessMessages)

	closeGrpcServer, grpcServer, err := s.newWriterGrpcServer()
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errcheck

	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errcheck

	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}

	s.runHealthCheck(ctx)
	s.runMetrics(cancel)

	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	<-ctx.Done()
	grpcServer.GracefulStop()

	return nil
}
