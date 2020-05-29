package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/josa42/mini-bcr/mentions-matcher/kafka"
	"github.com/josa42/mini-bcr/mentions-matcher/models"
)

var (
	broker         = env("KAFKA_HOST", "localhost:9092")
	subscribeTopic = env("KAFKA_SUBSCRIBE_TOPIC", "resources")
	publishTopic   = env("KAFKA_PUBLISH_TOPIC", "mentions")
)

func main() {
	kafkaClient := kafka.NewClient([]string{broker})

	kafkaClient.Subscribe(subscribeTopic, func(v string) {

		var resource models.Resource

		if err := json.Unmarshal([]byte(v), &resource); err == nil {
			mention := models.NewMentionFromResource(resource)

			kafkaClient.Publish(publishTopic, mention.ToJSON())

			log.Println(mention.ToJSON())
		}
	})
}

func env(key, defauleValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defauleValue
}
