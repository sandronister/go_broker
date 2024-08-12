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

	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", b.server, strconv.Itoa(b.port)),
		"group.id":          config["group.id"],
		"auto.offset.reset": config["auto.offset.reset"],
	}

	if config["auto.commit.enable"] != "" {
		kafkaConfig.SetKey("enable.auto.commit", true)
	}

	if config["auto.commit.interval.ms"] != "" {
		kafkaConfig.SetKey("auto.commit.interval.ms", config["auto.commit.interval.ms"])
	}

	return kafkaConfig

}
