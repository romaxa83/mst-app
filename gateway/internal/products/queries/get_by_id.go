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

type GetProductByIdHandler interface {
	Handle(ctx context.Context, query *GetProductByIdQuery) (*dto.ProductResponse, error)
}

type getProductByIdHandler struct {
	log      logger.Logger
	cfg      *config.Config
	rsClient readerService.ReaderServiceClient
}

func NewGetProductByIdHandler(log logger.Logger, cfg *config.Config, rsClient readerService.ReaderServiceClient) *getProductByIdHandler {
	return &getProductByIdHandler{log: log, cfg: cfg, rsClient: rsClient}
}

func (q *getProductByIdHandler) Handle(ctx context.Context, query *GetProductByIdQuery) (*dto.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getProductByIdHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())
	res, err := q.rsClient.GetProductById(ctx, &readerService.GetProductByIdReq{ProductID: query.ProductID.String()})
	if err != nil {
		return nil, err
	}

	return dto.ProductResponseFromGrpc(res.GetProduct()), nil
}
