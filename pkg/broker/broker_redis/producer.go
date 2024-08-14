package brokerredis

import (
	"context"
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Producer(message *types.Message, flush int) error {
	time.Sleep(time.Duration(flush) * time.Second)
	return b.client.Publish(context.Background(), message.TopicPartition, message.Value).Err()
}
