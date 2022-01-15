package service

import (
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/gateway/internal/author/commands"
	kafkaClient "github.com/romaxa83/mst-app/pkg/kafka"
	"github.com/romaxa83/mst-app/pkg/logger"
	readerService "github.com/romaxa83/mst-app/reader_service/proto/product_reader"
)

type AuthorService struct {
	Commands *commands.AuthorCmds
}

func NewAuthorService(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer, rsClient readerService.ReaderServiceClient) *AuthorService {

	createAuthorHandler := commands.NewCreateAuthorHandler(log, cfg, kafkaProducer)

	authorCommands := commands.NewAuthorCmds(createAuthorHandler)

	return &AuthorService{Commands: authorCommands}
}
