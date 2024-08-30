package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func read(info <-chan types.Message) {
	for msg := range info {
		fmt.Printf("recebida mensagem: %s\n", string(msg.Value))
	}

}

func main() {

	broker := factory.NewBroker(factory.REDIS, "localhost", "ruptela.com", 6379)

	if broker == nil {
		fmt.Println("erro ao criar broker")
		return
	}

	config := &types.ConfigBroker{
		Topic: "ruptela.com",
	}

	var ch = make(chan types.Message)

	for range 10 {
		go read(ch)
	}
	err := broker.Consumer(config, ch)
	if err != nil {
		fmt.Println("erro ao consumir mensagens:", err)
		return
	}

}
