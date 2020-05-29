package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/josa42/mini-bcr/pkg/kafka"
	"github.com/josa42/mini-bcr/pkg/models"
	"github.com/josa42/mini-bcr/twitter-crawler/twitter"
)

func main() {
	broker := env("KAFKA_HOST", "localhost:9092")
	topic := env("KAFKA_PUBLISH_TOPIC", "resources")

	consumerKey := env("TWITTER_CONSUMER_KEY", "")
	consumerSecret := env("TWITTER_CONSUMER_SECRET", "")
	token := env("TWITTER_TOKEN", "")
	tokenSecret := env("TWITTER_TOKEN_SECRET", "")

	if consumerKey == "" || consumerSecret == "" || token == "" || tokenSecret == "" {
		log.Println("Required: TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET, TWITTER_TOKEN and TWITTER_TOKEN_SECRET")
		os.Exit(1)
	}

	kafkaClient := kafka.NewClient([]string{broker})
	twitterClient := twitter.NewClient(consumerKey, consumerSecret, token, tokenSecret)

	twitterClient.Stream(func(m models.Resource) {
		go kafkaClient.Publish(topic, m.ToJSON())
		go logMention(m)
	})
}

func env(key, defauleValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defauleValue
}

func logMention(m models.Resource) {
	b, _ := json.Marshal(m)
	log.Println(string(b))
}
