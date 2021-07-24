package receiver

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/event"
)

var _ biz.OrderQueueReceiverRepo = (*orderQueueReceiverRepo)(nil)

type orderQueueReceiverRepo struct {
	queue *Receiver
	log *log.Helper
}

func (o orderQueueReceiverRepo) CreateOrder(ctx context.Context) error {

	o.queue.receiver.Receive(context.Background(), func(ctx context.Context, message event.Message) error {
		fmt.Printf("key:%s, value:%s, header:%s\n", message.Key(), message.Value(), message.Header())
		return nil
	})

	return nil
}


func NewOrderQueueReceiverRepo(queue *Receiver, logger log.Logger) biz.OrderQueueReceiverRepo {

	

	return &orderQueueReceiverRepo{
		queue: queue,
		log:  log.NewHelper(log.With(logger, "module", "queue/server-service")),
	}
}
