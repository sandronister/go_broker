package brokerredis

import (
	"context"
	"encoding/json"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Producer(message *types.Message) error {
	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return b.client.LPush(context.Background(), message.Topic, messageByte).Err()
}
