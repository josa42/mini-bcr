package twitter

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/josa42/mini-bcr/twitter-crawler/models"
)

type Client struct {
	client *twitter.Client
}

func NewClient(consumerKey, consumerSecret, token, tokenSecret string) *Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	otoken := oauth1.NewToken(token, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, otoken)

	client := twitter.NewClient(httpClient)

	return &Client{
		client: client,
	}
}

func (c *Client) Stream(cb func(resource models.Mention)) {

	stream, err := c.client.Streams.Sample(&twitter.StreamSampleParams{
		Language: []string{"de", "en"},
	})

	if err != nil {
		log.Fatal(err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		publishedAt, _ := tweet.CreatedAtTime()

		cb(models.Mention{
			URL:  fmt.Sprintf("https://twitter.com/%s/status/%d", tweet.User.ScreenName, tweet.ID),
			Text: tweet.Text,
			Author: models.Author{
				Name:     tweet.User.Name,
				Username: tweet.User.ScreenName,
			},
			PublishedAt: publishedAt,
			Source:      "twitter.com",
		})
	}

	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	stream.Stop()
}
