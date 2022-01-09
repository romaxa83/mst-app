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

type UpdateProductCmdHandler interface {
	Handle(ctx context.Context, command *UpdateProductCommand) error
}

type updateProductCmdHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewUpdateProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *updateProductCmdHandler {
	return &updateProductCmdHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *updateProductCmdHandler) Handle(ctx context.Context, command *UpdateProductCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateProductCmdHandler.Handle")
	defer span.Finish()

	updateDto := &kafkaMessages.ProductUpdate{
		ProductID:   command.UpdateDto.ProductID.String(),
		Name:        command.UpdateDto.Name,
		Description: command.UpdateDto.Description,
		Price:       command.UpdateDto.Price,
	}

	dtoBytes, err := proto.Marshal(updateDto)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.ProductUpdate.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
