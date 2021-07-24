package sender

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/kafka"
)

var _ biz.OrderQueueRepo = (*orderQueueRepo)(nil)

type orderQueueRepo struct {
	queue *Queue
	log *log.Helper
}

func (o orderQueueRepo) CreateOrder(ctx context.Context) error {

	msg := kafka.NewMessage("kratos", []byte("hello world"), map[string]string{
		"user": "kratos",
		"phone": "123123",
	})

	o.queue.sender.Send(ctx, msg)

	return nil
}

func NewOrderQueueRepo(queue *Queue, logger log.Logger) biz.OrderQueueRepo {

	return &orderQueueRepo{
		queue: queue,
		log:  log.NewHelper(log.With(logger, "module", "queue/server-service")),
	}
}
