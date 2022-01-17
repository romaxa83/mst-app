package kafka

import (
	"context"
	"github.com/romaxa83/mst-app/library-app/internal/config"
	//"github.com/romaxa83/mst-app/writer_service/internal/metrics"
	//"github.com/romaxa83/mst-app/writer_service/internal/product/service"
	"github.com/segmentio/kafka-go"
	"sync"
)

const (
	PoolSize = 30
)

type authorMessageProcessor struct {
	cfg *config.Config
	//ps      *service.ProductService
	//metrics *metrics.WriterServiceMetrics
}

func NewAuthorMessageProcessor(
	cfg *config.Config,
) *authorMessageProcessor {
	return &authorMessageProcessor{
		cfg: cfg,
		//ps: ps,
		//metrics: metrics,
	}
}

func (s *authorMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	//for {
	//	select {
	//	case <-ctx.Done():
	//		return
	//	default:
	//	}
	//
	//	m, err := r.FetchMessage(ctx)
	//	if err != nil {
	//		logger.Warnf("workerID: %v, err: %v", workerID, err)
	//		continue
	//	}
	//
	//	s.logProcessMessage(m, workerID)
	//
	//	switch m.Topic {
	//	case s.cfg.KafkaTopics.AuthorCreate.TopicName:
	//		s.processCreateAuthor(ctx, r, m)
	//	}
	//}
}
