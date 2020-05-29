package main

import (
	"context"
	"log"
	"os"

	"github.com/josa42/mini-bcr/pkg/kafka"
	"github.com/josa42/mini-bcr/pkg/models"
	"github.com/mattn/go-mastodon"
)

func main() {

	server := env("MASTODON_SERVER", "https://mas.to")
	clientID := env("MASTODON_CLIENT_ID", "")
	clientSecret := env("MASTODON_CLIENT_SECRET", "")
	username := env("MASTODON_USERNAME", "")
	password := env("MASTODON_PASSWORD", "")

	broker := env("KAFKA_HOST", "localhost:9092")
	topic := env("KAFKA_PUBLISH_TOPIC", "resources")

	kafkaClient := kafka.NewClient([]string{broker})

	// app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
	// 	Server:     server,
	// 	ClientName: "mini-bcr",
	// 	Scopes:     "read write follow",
	// 	Website:    "https://github.com/josa42/mini-bcr",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("client-id    : %s\n", app.ClientID)
	// fmt.Printf("client-secret: %s\n", app.ClientSecret)

	c := mastodon.NewClient(&mastodon.Config{
		Server:       server,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})

	err := c.Authenticate(context.Background(), username, password)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := c.StreamingPublic(context.Background(), false)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case event := <-stream:
			if update, ok := event.(*mastodon.UpdateEvent); ok {
				r := models.Resource{
					URL:  update.Status.URL,
					Text: update.Status.Content,
					Author: models.Author{
						Name:     update.Status.Account.DisplayName,
						Username: update.Status.Account.Acct,
					},
					PublishedAt: update.Status.CreatedAt,
					Source:      "mastodon",
				}

				kafkaClient.Publish(topic, r.ToJSON())
				logResource(r)
			}

		}
	}

}

func env(key, defauleValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defauleValue
}

func logResource(r models.Resource) {
	log.Println(r.ToJSON())
}

