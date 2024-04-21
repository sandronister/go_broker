package kafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go-broker/pkg"
)

func (b *Broker) Produce(message *pkg.Message) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", b.server, strconv.Itoa(b.port))})

	if err != nil {
		return err
	}

	topic := message.TopicPartition

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          []byte(message.Value),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte(message.Headers)}},
	}, nil)

	if err != nil {
		return err
	}

	p.Close()

	return nil
}
