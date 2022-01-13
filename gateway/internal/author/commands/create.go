package commands

import (
	"context"
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	kafkaClient "github.com/romaxa83/mst-app/pkg/kafka"
)

type CreateAuthorCmdHandler interface {
	Handle(ctx context.Context, cmd *CreateAuthorCmd) error
}

type CreateAuthorHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewCreateAuthorHandler(
	log logger.Logger,
	cfg *config.Config,
	kafkaProducer kafkaClient.Producer,
) *CreateAuthorHandler {
	return &CreateAuthorHandler{
		log:           log,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *CreateAuthorHandler) Handle(ctx context.Context, cmd *CreateAuthorCmd) error {

	logger.Warnf("CMD_AUTHOR v+%", cmd)

	return nil
}
