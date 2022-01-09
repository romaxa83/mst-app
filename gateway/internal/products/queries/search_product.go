package queries

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/gateway/internal/dto"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/pkg/tracing"
	readerService "github.com/romaxa83/mst-app/reader_service/proto/product_reader"
)

type SearchProductHandler interface {
	Handle(ctx context.Context, query *SearchProductQuery) (*dto.ProductsListResponse, error)
}

type searchProductHandler struct {
	log      logger.Logger
	cfg      *config.Config
	rsClient readerService.ReaderServiceClient
}

func NewSearchProductHandler(log logger.Logger, cfg *config.Config, rsClient readerService.ReaderServiceClient) *searchProductHandler {
	return &searchProductHandler{log: log, cfg: cfg, rsClient: rsClient}
}

func (s *searchProductHandler) Handle(ctx context.Context, query *SearchProductQuery) (*dto.ProductsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "searchProductHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())
	res, err := s.rsClient.SearchProduct(ctx, &readerService.SearchReq{
		Search: query.Text,
		Page:   int64(query.Pagination.GetPage()),
		Size:   int64(query.Pagination.GetSize()),
	})
	if err != nil {
		return nil, err
	}

	return dto.ProductsListResponseFromGrpc(res), nil
}
