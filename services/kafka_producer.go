package services

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var KafkaWriter = &kafka.Writer{
	Addr:     kafka.TCP("host.minikube.internal:9092"),
	Topic:    "audit-events",
	Balancer: &kafka.LeastBytes{},
}

func PublishAudit(message string) {

	err := KafkaWriter.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)

	if err != nil {
		log.Println("Kafka Publish Error:", err)
		return
	}

	log.Println("Kafka Message Published:", message)
}
