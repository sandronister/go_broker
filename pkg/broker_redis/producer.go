package brokerredis

import (
	"context"
	"time"

	"github.com/sandronister/go_broker/pkg/payload"
)

func (b *Broker) Produce(message *payload.Message, flush int) error {
	time.Sleep(time.Duration(flush) * time.Second)
	return b.client.Publish(context.Background(), message.TopicPartition, message.Value).Err()
}
