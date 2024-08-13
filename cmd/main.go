package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	brokerredis "github.com/sandronister/go_broker/pkg/broker_redis"
	"github.com/sandronister/go_broker/pkg/payload"
)

func printMessage(message <-chan payload.Message, db *sql.DB) {
	for msg := range message {

		_, err := db.Exec("INSERT INTO messages (message) VALUES (?)", msg.Value)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		fmt.Println("Successfully inserted message", string(msg.Value))

	}
}

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

	ch, err := broker.Consumer("myTopic")
	if err != nil {
		fmt.Println("Erro ao consumir mensagens:", err)
		return
	}

	for msg := range ch {
		fmt.Printf("Recebida mensagem: %s\n", string(msg.Value))
	}
}
