package Kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

func createProducer(broker string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatal("Failed to start Sarama producer:", err)
	}
	return producer
}
