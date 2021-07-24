package receiver

import (
	"github.com/google/wire"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/event"
	"github.com/peter-wow/seckill/app/order/service/internal/queue/kafka"
	"log"
)

var ProviderSet = wire.NewSet(NewQueue, NewOrderQueueReceiverRepo)

type Receiver struct {
	receiver event.Receiver
}

func NewQueue() *Receiver {

	log.Println("queue init")

	receiverClient, err := kafka.NewKafkaReceiver([]string{"192.168.2.27:9092"}, "test")

	if err != nil {
		log.Fatalln("kafka erceiver fatal: ", err)
	}

	return &Receiver{
		receiver: receiverClient,
	}
}