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

		for _, topic := range config.Topic {
			err := c.Assign([]kafka.TopicPartition{
				{Topic: &topic, Partition: int32(config.Partition)},
			})

			if err != nil {
				return nil, err
			}
		}
	}

	if config.Partition == 0 {
		err := c.SubscribeTopics(config.Topic, nil)

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
