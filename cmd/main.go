package main

import (
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func main() {

	broker := factory.NewBroker(factory.REDIS, "localhost", "teste", 6379)

	if broker == nil {
		fmt.Println("erro ao criar broker")
		return
	}

	config := &types.ConfigBroker{
		Topic: "ruptela.com",
	}

	var ch = make(chan types.Message)
	err := broker.Consumer(config, ch)
	if err != nil {
		fmt.Println("erro ao consumir mensagens:", err)
		return
	}

	i := 0
	start := time.Now()
	for msg := range ch {
		fmt.Printf("recebida mensagem: %s\n", string(msg.Value))
		i++
		fmt.Printf("total de mensagens recebidas: %d\n", i)
		if i == 333 {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("tempo total: %s\n", elapsed)
}
