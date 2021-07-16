package queue

import(
	"github.com/Shopify/sarama"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewQueue, NewOrderQueueRepo)

type Queue struct {
	producer sarama.AsyncProducer
}

func NewQueue() *Queue {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	client,err := sarama.NewClient([]string{"192.168.2.27:9092"}, config)



	defer client.Close()
	if err != nil {
		panic(err)
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)

	queue := &Queue{
		producer: producer,
	}

	return queue
}


