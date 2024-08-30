package brokerkafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) GetConfig(config *types.ConfigBroker) *kafka.ConfigMap {
	if config.GroupName == "" {
		config.GroupName = "default"
	}

	if config.AutoOffsetReset == "" {
		config.AutoOffsetReset = "earliest"
	}

	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", b.host, strconv.Itoa(b.port)),
		"group.id":          config.GroupName,
		"auto.offset.reset": config.AutoOffsetReset,
	}

	if !config.EnableAutoCommit {
		kafkaConfig.SetKey("enable.auto.commit", true)
	}

	if config.AutoCommitIntervalMS != "" {
		kafkaConfig.SetKey("auto.commit.interval.ms", config.AutoCommitIntervalMS)
	}

	return kafkaConfig

}
