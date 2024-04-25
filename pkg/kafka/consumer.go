package kafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go-broker/pkg/payload"
)

func (b *Broker) Consume(topic string, group string, message chan<- payload.Message) error {
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

		var listHeaders []payload.Header

		for _, h := range msg.Headers {
			header := payload.Header{Key: h.Key, Value: string(h.Value)}
			listHeaders = append(listHeaders, header)
		}

		message <- payload.Message{
			TopicPartition: *msg.TopicPartition.Topic,
			Value:          msg.Value,
			Key:            msg.Key,
			Timestamp:      msg.Timestamp,
			Headers:        listHeaders,
		}

	}

}
