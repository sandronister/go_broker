package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/sandronister/go-broker/pkg"
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

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)

	message := make(chan payload.Message)

	for range 10 {
		go printMessage(message)
	}

	for i := range 3 {
		go broker.Consume(pkg.ConfigMap{
			"topic":             "new-topic",
			"group.id":          "my-group",
			"auto.offset.reset": "earliest",
			"partition":         strconv.Itoa(i),
		}, message)

	}

	waitGroup.Wait()
}
