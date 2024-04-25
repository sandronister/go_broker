package main

import (
	"fmt"

	"github.com/sandronister/go-broker/pkg/kafka"
	"github.com/sandronister/go-broker/pkg/payload"
)

func printMessage(message <-chan payload.Message) {
	for msg := range message {
		println("Message Received:")
		println("TopicPartition: ", msg.TopicPartition)
		println("Value: ", string(msg.Value))
		println("Key: ", string(msg.Key))

		for _, h := range msg.Headers {
			fmt.Printf("Headers: Key %s, Value %s\n", h.Key, h.Value)
		}
	}
}

func main() {

	broker := kafka.NewBroker("localhost", 9092)

	message := make(chan payload.Message)

	for i := 0; i < 10; i++ {
		go printMessage(message)
	}

	err := broker.Consume("myTopic", "myGroup", message)

	if err != nil {
		panic(err)
	}

}
