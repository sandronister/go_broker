package kafka

import (
	"github.com/sandronister/go-broker/pkg/connection"
	"github.com/sandronister/go-broker/pkg/payload"
)

func (b *Broker) Consume(config connection.ConfigMap, message chan<- payload.Message) error {

	c, err := b.getConsumer(config)

	if err != nil {
		return err
	}

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
