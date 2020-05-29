package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Encodable interface {
	Encode() string
}

type Client struct {
	brokers []string
}

func NewClient(brokers []string) *Client {
	return &Client{brokers: brokers}
}

func (c *Client) Send(topic string, value Encodable) {
	producer, err := sarama.NewSyncProducer(c.brokers, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(value.Encode())}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	}

}
