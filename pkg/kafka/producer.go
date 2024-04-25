package kafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go-broker/pkg/payload"
)

func (b *Broker) Produce(message *payload.Message) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", b.server, strconv.Itoa(b.port))})

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

	p.Close()

	return nil
}
