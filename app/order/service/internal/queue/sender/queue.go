package sender

import (
	"github.com/google/wire"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/event"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/kafka"
	"log"
)

var ProviderSet = wire.NewSet(NewQueue, NewOrderQueueRepo)

type Queue struct {
	sender event.Sender
}

func NewQueue() *Queue {

	log.Println("queue init")

	senderClient, err := kafka.NewKafkaSender([]string{"192.168.2.27:9092"}, "test")

	if err != nil {
		log.Fatalln("kafka sender client fatal: ", err)
	}

	return &Queue{
		sender: senderClient,
	}
}