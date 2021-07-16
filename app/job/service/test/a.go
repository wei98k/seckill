package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {

	

	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	client,err := sarama.NewClient([]string{"192.168.2.27:9092"}, config)
	defer client.Close()
	if err != nil {
		panic(err)
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg        sync.WaitGroup
		enqueued, successes, errors int
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

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}
