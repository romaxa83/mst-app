package kafka

import (
	"context"
	"github.com/avast/retry-go"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"github.com/romaxa83/mst-app/pkg/tracing"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (s *authorMessageProcessor) processCreateAuthor(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	//s.metrics.CreateProductKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "authorMessageProcessor.processCreateAuthor")
	defer span.Finish()

	//var msg kafkaMessages.AuthorCreate
	logger.Warn(m)
	//if err := proto.Unmarshal(m.Value, &msg); err != nil {
	//	s.log.WarnMsg("proto.Unmarshal", err)
	//	s.commitErrMessage(ctx, r, m)
	//	return
	//}
	//
	//proUUID, err := uuid.FromString(msg.GetProductID())
	//if err != nil {
	//	s.log.WarnMsg("proto.Unmarshal", err)
	//	s.commitErrMessage(ctx, r, m)
	//	return
	//}
	//
	//command := commands.NewCreateProductCommand(proUUID, msg.GetName(), msg.GetDescription(), msg.GetPrice())
	//if err := s.v.StructCtx(ctx, command); err != nil {
	//	s.log.WarnMsg("validate", err)
	//	s.commitErrMessage(ctx, r, m)
	//	return
	//}
	//
	//if err := retry.Do(func() error {
	//	return s.ps.Commands.CreateProduct.Handle(ctx, command)
	//}, append(retryOptions, retry.Context(ctx))...); err != nil {
	//	s.log.WarnMsg("CreateProduct.Handle", err)
	//	s.metrics.ErrorKafkaMessages.Inc()
	//	return
	//}

	s.commitMessage(ctx, r, m)
}
