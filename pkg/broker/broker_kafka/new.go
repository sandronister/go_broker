package brokerkafka

import "github.com/sandronister/go_broker/pkg/broker/types"

type Broker struct {
	host  string
	port  int
	topic string
}

func NewBroker(host string, topic string, port int) types.IBroker {
	return &Broker{
		host:  host,
		port:  port,
		topic: topic,
	}
}
