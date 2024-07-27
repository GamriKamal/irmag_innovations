package Kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

var counter int

func produceMessages(producer sarama.AsyncProducer, topic string, messages []map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, message := range messages {

		jsonMessage, err := json.Marshal(message)
		if err != nil {
			log.Println("Error marshalling message:", err)
			continue
		}

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(jsonMessage),
		}

		producer.Input() <- msg
		counter++

		select {
		case <-producer.Successes():
			log.Println("Message sent successfully", topic, counter)
		case <-producer.Errors():
			log.Println("Failed to send message", topic)
		}
	}
}
