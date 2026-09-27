package services

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"host.minikube.internal:9092"},
		Topic:   "audit-events",
		GroupID: "audit-group",
	})

	log.Println("Kafka Consumer Started...")

	for {

		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Consumer Error:", err)
			continue
		}

		log.Println("Received:", string(msg.Value))
	}
}
