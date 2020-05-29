package main

import (
	"log"
	"os"

	"github.com/josa42/mini-bcr/mentions-storage/elastic"
	"github.com/josa42/mini-bcr/pkg/kafka"
	"github.com/josa42/mini-bcr/pkg/models"
)

var (
	elasticHost    = env("ELASTICSEARCH_HOST", "localhost:9200")
	broker         = env("KAFKA_HOST", "localhost:9092")
	subscribeTopic = env("KAFKA_SUBSCRIBE_TOPIC", "mentions")
)

func main() {
	log.Printf("broker: %s", broker)
	log.Printf("elastic: %s", elasticHost)

	elasticClient := elastic.NewClient(elasticHost)
	kafkaClient := kafka.NewClient([]string{broker})

	log.Println(elasticClient.Info())

	kafkaClient.Subscribe(subscribeTopic, func(value string) {
		m := models.NewMentionFromJSON(value)
		elasticClient.Store(m)
	})
}

func env(key, defauleValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defauleValue
}

