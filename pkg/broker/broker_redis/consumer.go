package brokerredis

import (
	"context"
	"fmt"
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Consumer(config *types.ConfigBroker, message chan<- types.Message) error {
	if config == nil {
		return types.ErrInvalidConfig
	}

	if config.Topic == "" {
		return types.ErrInvalidConfig
	}

	for {
		res, err := b.client.BLPop(context.Background(), 0*time.Second, config.Topic).Result()
		if err != nil {
			fmt.Println("Erro ao ler item da fila:", err)
			continue
		}

		message <- types.Message{
			Value: []byte(res[1]),
		}
	}

}
