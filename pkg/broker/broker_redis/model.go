package brokerredis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

type Broker struct {
	client *redis.Client
}

func NewBroker(server string, port int) types.IBroker {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
	})
	return &Broker{
		client: client,
	}
}
