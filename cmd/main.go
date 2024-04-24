package main

import (
	"github.com/sandronister/go-broker/pkg"
	"github.com/sandronister/go-broker/pkg/kafka"
)

func printMessage(message <-chan pkg.ReceiptMessage) {
	for msg := range message {
		println("Message Received:")
		println("TopicPartition: ", msg.TopicPartition)
		println("Value: ", string(msg.Value))
		println("Key: ", string(msg.Key))

		for _, h := range msg.Headers {
			println("Headers: ", h)
		}
	}
}

func main() {

	broker := kafka.NewBroker("localhost", 9092)

	message := make(chan pkg.ReceiptMessage)

	for i := 0; i < 10; i++ {
		go printMessage(message)
	}

	err := broker.Consume("myTopic", message)

	if err != nil {
		println(err)
	}

}
