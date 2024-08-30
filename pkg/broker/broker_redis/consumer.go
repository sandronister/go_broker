package brokerredis

import (
	"context"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Consumer(config *types.ConfigBroker, message chan<- types.Message) error {
	if config == nil {
		return types.ErrInvalidConfig
	}

	if config.Topic == "" {
		return types.ErrInvalidConfig
	}

	pubsub := b.client.Subscribe(context.Background(), config.Topic)
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
