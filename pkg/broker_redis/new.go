package brokerredis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sandronister/go_broker/pkg/ports"
)

type Broker struct {
	client *redis.Client
}

func NewBroker(server string, port int) ports.IBroker {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
		DB:   0,
	})
	return &Broker{
		client: client,
	}
}
