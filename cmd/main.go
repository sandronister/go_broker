package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func read(info <-chan types.Message) {
	for msg := range info {
		fmt.Printf("recebida mensagem: %s no topico %s \n", string(msg.Value), string(msg.Topic))
	}

}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

}

func main() {

	fmt.Println("iniciando broker")
	broker, err := factory.GetBroker()

	if err != nil {
		fmt.Println("erro ao criar broker", err)
		return
	}

	config := &types.ConfigBroker{
		Topic: []string{"page"},
	}

	var ch = make(chan types.Message)

	for range 10 {
		go read(ch)
	}

	err = broker.ListenToQueue(config, ch)
	if err != nil {
		fmt.Println("erro ao consumir mensagens:", err)
		return
	}

	fmt.Println("esperando mensagens")

}
