package kafka

import (
	"context"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"github.com/segmentio/kafka-go"
)

func (s *authorMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	//s.metrics.SuccessKafkaMessages.Inc()
	logger.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		logger.Warn("commitMessage", err)
	}
}

func (s *authorMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	//s.metrics.ErrorKafkaMessages.Inc()
	logger.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		logger.Warn("commitMessage", err)
	}
}

func (s *authorMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	logger.KafkaProcessMessage(m.Topic, m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}
