package brokerkafka

import "github.com/sandronister/go_broker/pkg/broker/types"

type Broker struct {
	host string
	port int
}

func NewBroker(host string, port int) types.IBroker {
	return &Broker{
		host: host,
		port: port,
	}
}
