package brokerkafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) getConsumer(config *types.ConfigBroker) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(b.GetConfig(config))

	if err != nil {
		return nil, err
	}

	if config.Partition != 0 {

		c.Assign([]kafka.TopicPartition{{Topic: &config.Topic, Partition: int32(config.Partition)}})
	}

	if config.Partition == 0 {
		err := c.SubscribeTopics([]string{config.Topic}, nil)

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
