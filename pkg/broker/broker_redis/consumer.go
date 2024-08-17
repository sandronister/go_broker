package brokerredis

import (
	"context"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Consumer(config *types.ConfigMap, message chan<- types.Message) error {
	if config == nil {
		return types.ErrInvalidConfig
	}

	topic, ok := (*config)["topic"]

	if !ok {
		return types.ErrInvalidConfig
	}

	pubsub := b.client.Subscribe(context.Background(), topic)
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			message <- types.Message{
				Value: []byte(msg.Payload),
			}
		}
	}()

	return nil
}
