package brokerredis

import (
	"context"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Producer(message *types.Message) error {
	return b.client.Publish(context.Background(), b.Topic, message.Value).Err()
}
