package kafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go-broker/pkg"
)

func (b *Broker) Consume(topic string, group string, message chan<- pkg.ReceiptMessage) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", b.server, strconv.Itoa(b.port)),
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return err
	}

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			return err
		}

		var headers []string

		for _, h := range msg.Headers {
			headers = append(headers, fmt.Sprintf("%s: %s", h.Key, string(h.Value)))
		}

		message <- pkg.ReceiptMessage{
			TopicPartition: *msg.TopicPartition.Topic,
			Value:          msg.Value,
			Key:            msg.Key,
			Timestamp:      msg.Timestamp,
			Headers:        headers,
		}

	}

}
