package kafka

import (
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg"
)

func (b *Broker) getConsumer(config pkg.ConfigMap) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(b.GetConfig(config))

	if err != nil {
		return nil, err
	}

	if config["partition"] != "" {
		topic := config["topic"]
		partition, err := strconv.Atoi(config["partition"])

		if err != nil {
			return nil, err
		}
		c.Assign([]kafka.TopicPartition{{Topic: &topic, Partition: int32(partition)}})
	}

	if config["partition"] == "" {
		err := c.SubscribeTopics([]string{config["topic"]}, nil)

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
