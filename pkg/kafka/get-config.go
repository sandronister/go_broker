package kafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (b *Broker) GetConfig(config ConfigMap) *kafka.ConfigMap {
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

	return kafkaConfig

}
