package brokerkafka

import (
	"github.com/sandronister/go_broker/pkg/payload"
	"github.com/sandronister/go_broker/pkg/ports"
)

func (b *Broker) Consume(config ports.ConfigMap, message chan<- payload.Message) error {

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
