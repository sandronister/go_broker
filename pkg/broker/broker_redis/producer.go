package brokerredis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Producer(message *types.Message) error {
	err := b.client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: message.GroupID,
		Values: message.Value,
	}).Err()

	if err != nil {
		return err
	}

	return nil
}
