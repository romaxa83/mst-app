package queries

import (
	"context"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/reader_service/config"
	"github.com/romaxa83/mst-app/reader_service/internal/models"
	"github.com/romaxa83/mst-app/reader_service/internal/product/repository"
)

type SearchProductHandler interface {
	Handle(ctx context.Context, query *SearchProductQuery) (*models.ProductsList, error)
}

type searchProductHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewSearchProductHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *searchProductHandler {
	return &searchProductHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (s *searchProductHandler) Handle(ctx context.Context, query *SearchProductQuery) (*models.ProductsList, error) {
	return s.mongoRepo.Search(ctx, query.Text, query.Pagination)
}
