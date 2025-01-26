package brokerredis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Publish(message *types.Message) error {
	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return b.client.LPush(context.Background(), message.Topic, messageByte).Err()
}

func (b *Broker) ListenToQueue(config *types.ConfigBroker, message chan<- types.Message) error {
	if config == nil {
		return types.ErrInvalidConfig
	}

	if config.Topic == nil {
		return types.ErrInvalidConfig
	}

	for {
		res, err := b.client.BLPop(context.Background(), 0*time.Second, config.Topic...).Result()
		if err != nil {
			fmt.Println("Erro ao ler item da fila:", err)
			continue
		}

		var tmp types.Message

		json.Unmarshal([]byte(res[1]), &tmp)

		message <- tmp

	}
}
