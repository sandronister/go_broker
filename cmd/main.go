package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/sandronister/go-broker/pkg/kafka"
	"github.com/sandronister/go-broker/pkg/payload"
)

func printMessage(message <-chan payload.Message, inx int) {
	for msg := range message {
		println("Message Received:")
		println("TopicPartition: ", msg.TopicPartition)
		println("Value: ", string(msg.Value))
		println("Key: ", string(msg.Key))

		fmt.Printf("Index %s\n", strconv.Itoa(inx))

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

	for i := 0; i < 10; i++ {
		go printMessage(message, i)
	}

	for i := range 3 {
		go broker.Consume(kafka.ConfigMap{
			"topic":             "new-topic",
			"group.id":          "my-group",
			"auto.offset.reset": "earliest",
			"partition":         strconv.Itoa(i),
		}, message)

	}

	waitGroup.Wait()
}
