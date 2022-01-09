package service

import (
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/gateway/internal/products/commands"
	"github.com/romaxa83/mst-app/gateway/internal/products/queries"
	kafkaClient "github.com/romaxa83/mst-app/pkg/kafka"
	"github.com/romaxa83/mst-app/pkg/logger"
	readerService "github.com/romaxa83/mst-app/reader_service/proto/product_reader"
)

type ProductService struct {
	Commands *commands.ProductCommands
	Queries  *queries.ProductQueries
}

func NewProductService(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer, rsClient readerService.ReaderServiceClient) *ProductService {

	createProductHandler := commands.NewCreateProductHandler(log, cfg, kafkaProducer)
	updateProductHandler := commands.NewUpdateProductHandler(log, cfg, kafkaProducer)
	deleteProductHandler := commands.NewDeleteProductHandler(log, cfg, kafkaProducer)

	getProductByIdHandler := queries.NewGetProductByIdHandler(log, cfg, rsClient)
	searchProductHandler := queries.NewSearchProductHandler(log, cfg, rsClient)

	productCommands := commands.NewProductCommands(createProductHandler, updateProductHandler, deleteProductHandler)
	productQueries := queries.NewProductQueries(getProductByIdHandler, searchProductHandler)

	return &ProductService{Commands: productCommands, Queries: productQueries}
}
