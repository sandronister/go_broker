package brokerkafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) Producer(message *types.Message, flush int) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", b.host, strconv.Itoa(b.port))})

	if err != nil {
		return err
	}

	topic := message.TopicPartition

	headers := b.getHeader(message)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          []byte(message.Value),
		Headers:        headers,
	}, nil)

	if err != nil {
		return err
	}

	p.Flush(flush)

	p.Close()

	return nil
}
