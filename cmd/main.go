package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func read(info <-chan types.Message) {
	for msg := range info {
		fmt.Printf("recebida mensagem: %s no topico %s \n", string(msg.Value), string(msg.Topic))
	}

}

func main() {

	broker := factory.NewBroker(factory.REDIS, "localhost", 6379)

	if broker == nil {
		fmt.Println("erro ao criar broker")
		return
	}

	config := &types.ConfigBroker{
		Topic: []string{"bananinha", "abobrinha"},
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
	fmt.Println("esperando mensagens")

}
