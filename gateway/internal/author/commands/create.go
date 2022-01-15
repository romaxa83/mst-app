package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/romaxa83/mst-app/gateway/config"
	kafkaClient "github.com/romaxa83/mst-app/pkg/kafka"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/pkg/tracing"
	kafkaMessages "github.com/romaxa83/mst-app/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
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

	span, ctx := opentracing.StartSpanFromContext(ctx, "createAuthorHandler.Handle")
	defer span.Finish()

	createDto := &kafkaMessages.AuthorCreate{
		ID:       cmd.CreateDto.ID.String(),
		Name:     cmd.CreateDto.Name,
		Bio:      cmd.CreateDto.Bio,
		Birthday: cmd.CreateDto.Birthday.Unix(),
	}

	dtoBytes, err := proto.Marshal(createDto)
	if err != nil {
		return err
	}
	c.log.Warnf("CMD_AUTHOR %+v", dtoBytes)

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.AuthorCreate.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
