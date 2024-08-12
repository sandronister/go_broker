package kafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/ports"
)

func (b *Broker) GetConfig(config ports.ConfigMap) *kafka.ConfigMap {
	if config["group.id"] == "" {
		config["group.id"] = "default"
	}

	if config["auto.offset.reset"] == "" {
		config["auto.offset.reset"] = "earliest"
	}

	if config["enable.auto.commit"] == "" {
		config["enable.auto.commit"] = "true"
	}

	if config["auto.commit.interval.ms"] == "" {
		config["auto.commit.interval.ms"] = "1000"
	}

	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", b.server, strconv.Itoa(b.port)),
		"group.id":          config["group.id"],
		"auto.offset.reset": config["auto.offset.reset"],
	}

	return kafkaConfig

}
