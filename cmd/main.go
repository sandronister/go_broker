package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	brokerredis "github.com/sandronister/go_broker/pkg/broker_redis"
	"github.com/sandronister/go_broker/pkg/payload"
)

func main() {

	broker := brokerredis.NewBroker("localhost", 6379)

	go func() {
		for {
			msg := &payload.Message{
				TopicPartition: "myTopic",
				Value:          []byte("Hello, Redis!"),
			}
			err := broker.Produce(msg, 2)
			if err != nil {
				fmt.Println("Erro ao produzir mensagem:", err)
			}

		}
	}()

	var ch = make(chan payload.Message)
	err := broker.Consume(map[string]string{"topic": "myTopic"}, ch)
	if err != nil {
		fmt.Println("Erro ao consumir mensagens:", err)
		return
	}

	for msg := range ch {
		fmt.Printf("Recebida mensagem: %s\n", string(msg.Value))
	}
}
