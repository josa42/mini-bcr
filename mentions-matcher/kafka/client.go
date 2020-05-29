package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Client struct {
	brokers []string
}

func NewClient(brokers []string) *Client {
	return &Client{brokers: brokers}
}

func (c *Client) Publish(topic string, value string) {
	producer, err := sarama.NewSyncProducer(c.brokers, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(value)}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	}
}

func (c *Client) Subscribe(topic string, cb func(value string)) func() {

	consumer, err := sarama.NewConsumer(c.brokers, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			cb(string(msg.Value))
		}
	}
}

