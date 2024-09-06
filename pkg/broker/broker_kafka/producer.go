package brokerkafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Producer(message *types.Message) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", b.host, strconv.Itoa(b.port))})

	if err != nil {
		return err
	}

	headers := b.getHeader(message)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &message.Topic, Partition: 0},
		Value:          []byte(message.Value),
		Headers:        headers,
	}, nil)

	if err != nil {
		return err
	}

	p.Close()

	return nil
}
