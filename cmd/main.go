package main

import (
	"github.com/sandronister/go-broker/pkg"
	"github.com/sandronister/go-broker/pkg/kafka"
)

func main() {

	broker := kafka.NewBroker("localhost", 9092)

	message := &pkg.Message{
		TopicPartition: "myTopic",
		Value:          "New Producer Hello Go!",
		Headers:        "header values are binary",
	}

	broker.Produce(message)

}
