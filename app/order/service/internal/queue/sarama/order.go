package sarama

import (
	"context"
	"encoding/json"
	"fmt"
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
	// o.queue.producer.Input()

	o.log.Info("queue info, %d", o.queue.producer)

	topic := "test"
	value := "are you ok"

	data, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("[kafka_producer][sendMessage]:%s", err.Error())
		//g.Log.Error("[kafka_producer][sendMessage]:%s", err.Error())
		return nil
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	o.queue.producer.Input() <- msg


	return nil
}

func NewOrderQueueRepo(queue *Queue, logger log.Logger) biz.OrderQueueRepo {
	return &orderQueueRepo{
		queue: queue,
		log:  log.NewHelper(log.With(logger, "module", "queue/server-service")),
	}
}
