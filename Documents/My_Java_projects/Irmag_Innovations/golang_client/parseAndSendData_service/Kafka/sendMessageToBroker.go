package Kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

func SendMessageToBroker(dataChannel chan map[string]interface{}, category string) {
	var data []map[string]interface{}

	var wg sync.WaitGroup

	for msg := range dataChannel {
		data = append(data, msg)
	}

	wg.Add(1)

	go func(category string, data []map[string]interface{}) {
		defer wg.Done()

		producer := createProducer("localhost:29092")
		defer producer.AsyncClose()

		wg.Add(2)
		go handleErrors(producer, &wg)

		topic := category
		produceMessages(producer, topic, data, &wg)
	}(category, data)
	wg.Wait()

}

func handleErrors(producer sarama.AsyncProducer, wg *sync.WaitGroup) {
	defer wg.Done()
	for err := range producer.Errors() {
		log.Println("Failed to produce message:", err)
	}
}
