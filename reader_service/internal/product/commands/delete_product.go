package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/reader_service/config"
	"github.com/romaxa83/mst-app/reader_service/internal/product/repository"
)

type DeleteProductCmdHandler interface {
	Handle(ctx context.Context, command *DeleteProductCommand) error
}

type deleteProductCmdHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewDeleteProductCmdHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *deleteProductCmdHandler {
	return &deleteProductCmdHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (c *deleteProductCmdHandler) Handle(ctx context.Context, command *DeleteProductCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteProductCmdHandler.Handle")
	defer span.Finish()

	if err := c.mongoRepo.DeleteProduct(ctx, command.ProductID); err != nil {
		return err
	}

	c.redisRepo.DelProduct(ctx, command.ProductID.String())
	return nil
}
