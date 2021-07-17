package queue

import(
	"github.com/Shopify/sarama"
	"github.com/google/wire"
	"log"
	"sync"
)

var ProviderSet = wire.NewSet(NewQueue, NewOrderQueueRepo)

type Queue struct {
	producer sarama.AsyncProducer
}

func NewQueue() *Queue {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	//client,err := sarama.NewClient([]string{"192.168.2.27:9092"}, config)
	//
	//defer client.Close()

	//if err != nil {
	//	panic(err)
	//}

	producer, err := sarama.NewAsyncProducer([]string{"192.168.2.27:9092"}, config)

	if err != nil {
		panic(err)
	}

	queue := &Queue{
		producer: producer,
	}

	var (
		wg        sync.WaitGroup
		successes, errors int
	)

	wg.Add(1)
	// start a groutines to count successes num
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	// start a groutines to count error num
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()


	return queue
}


