package queue

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)

var _ biz.OrderQueueRepo = (*orderQueueRepo)(nil)

type orderQueueRepo struct {
	queue *Queue
	log *log.Helper
}

func (o orderQueueRepo) CreateOrder(ctx context.Context) error {
	// Trap SIGINT to trigger a graceful shutdown.

	o.log.Info("queue info")

	message := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("testing 123")}

	o.queue.producer.Input() <- message



	o.log.Infof("Successfully produced: %d; ", message)

	return nil
	panic("implement me")
}

func NewOrderQueueRepo(queue *Queue, logger log.Logger) biz.OrderQueueRepo {
	return &orderQueueRepo{
		queue: queue,
		log:  log.NewHelper(log.With(logger, "module", "queue/server-service")),
	}
}