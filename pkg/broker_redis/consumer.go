package brokerredis

import (
	"context"

	"github.com/sandronister/go_broker/pkg/payload"
	"github.com/sandronister/go_broker/pkg/ports"
)

func (b *Broker) Consume(config ports.ConfigMap, message chan<- payload.Message) error {
	if config == nil {
		return ports.ErrInvalidConfig
	}

	topic, ok := config["topic"]

	if !ok {
		return ports.ErrInvalidConfig
	}

	pubsub := b.client.Subscribe(context.Background(), topic)
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			message <- payload.Message{
				TopicPartition: msg.Channel,
				Value:          []byte(msg.Payload),
			}
		}
	}()

	return nil
}
