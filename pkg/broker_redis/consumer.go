package brokerredis

import (
	"context"

	"github.com/sandronister/go_broker/pkg/payload"
)

func (b *Broker) Consumer(topic string) (<-chan payload.Message, error) {
	pubsub := b.client.Subscribe(context.Background(), topic)
	ch := pubsub.Channel()
	message := make(chan payload.Message)

	go func() {
		for msg := range ch {
			message <- payload.Message{
				TopicPartition: msg.Channel,
				Value:          []byte(msg.Payload),
			}
		}
	}()

	return message, nil
}
