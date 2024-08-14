package brokerkafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) getHeader(message *types.Message) []kafka.Header {
	var listHeader []kafka.Header

	for _, item := range message.Headers {
		header := kafka.Header{Key: item.Key, Value: []byte(item.Value)}
		listHeader = append(listHeader, header)
	}

	return listHeader
}
