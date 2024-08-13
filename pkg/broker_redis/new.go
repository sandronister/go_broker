package brokerredis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Broker struct {
	client *redis.Client
}

func NewBroker(server string, port int) *Broker {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
		DB:   0,
	})
	return &Broker{
		client: client,
	}
}
