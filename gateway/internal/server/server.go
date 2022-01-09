package server

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/gateway/internal/client"
	"github.com/romaxa83/mst-app/gateway/internal/metrics"
	"github.com/romaxa83/mst-app/gateway/internal/middlewares"
	v1 "github.com/romaxa83/mst-app/gateway/internal/products/delivery/http/v1"
	"github.com/romaxa83/mst-app/gateway/internal/products/service"
	"github.com/romaxa83/mst-app/pkg/interceptors"
	"github.com/romaxa83/mst-app/pkg/kafka"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/pkg/tracing"
	readerService "github.com/romaxa83/mst-app/reader_service/proto/product_reader"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log  logger.Logger
	cfg  *config.Config
	v    *validator.Validate
	mw   middlewares.MiddlewareManager
	im   interceptors.InterceptorManager
	echo *echo.Echo
	ps   *service.ProductService
	m    *metrics.ApiGatewayMetrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, echo: echo.New(), v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.mw = middlewares.NewMiddlewareManager(s.log, s.cfg)
	s.im = interceptors.NewInterceptorManager(s.log)
	s.m = metrics.NewApiGatewayMetrics(s.cfg)

	readerServiceConn, err := client.NewReaderServiceConn(ctx, s.cfg, s.im)
	if err != nil {
		return err
	}
	defer readerServiceConn.Close() // nolint: errcheck
	rsClient := readerService.NewReaderServiceClient(readerServiceConn)

	kafkaProducer := kafka.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errcheck

	s.ps = service.NewProductService(s.log, s.cfg, kafkaProducer, rsClient)

	productHandlers := v1.NewProductsHandlers(s.echo.Group(s.cfg.Http.ProductsPath), s.log, s.mw, s.cfg, s.ps, s.v, s.m)
	productHandlers.MapRoutes()

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)

	s.runMetrics(cancel)
	s.runHealthCheck(ctx)

	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	<-ctx.Done()
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}

	return nil
}
