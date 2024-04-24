package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go-broker/pkg"
)

func (b *Broker) getHeader(message *pkg.SendMessage) []kafka.Header {
	var listHeader []kafka.Header

	for _, item := range message.Headers {
		header := kafka.Header{Key: item.Key, Value: []byte(item.Value)}
		listHeader = append(listHeader, header)
	}

	return listHeader
}
