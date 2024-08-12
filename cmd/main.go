package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/go_broker/pkg/kafka"
	"github.com/sandronister/go_broker/pkg/payload"
	"github.com/sandronister/go_broker/pkg/ports"
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

	broker := kafka.NewBroker("localhost", 9092)

	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)

	message := make(chan payload.Message)
	message2 := make(chan payload.Message)
	message3 := make(chan payload.Message)

	for range 100 {
		go printMessage(message, db)
		go printMessage(message2, db)
		go printMessage(message3, db)
	}

	go broker.Consume(ports.ConfigMap{
		"topic":                   "omnicom.com",
		"group.id":                "my-group",
		"auto.offset.reset":       "earliest",
		"partition":               "0",
		"auto.commit.enable":      "true",
		"auto.commit.interval.ms": "1000",
	}, message)

	go broker.Consume(ports.ConfigMap{
		"topic":                   "ruptela.com",
		"group.id":                "my-group",
		"auto.offset.reset":       "earliest",
		"partition":               "0",
		"auto.commit.enable":      "true",
		"auto.commit.interval.ms": "1000",
	}, message2)

	go broker.Consume(ports.ConfigMap{
		"topic":                   "sinocastle.com",
		"group.id":                "my-group",
		"auto.offset.reset":       "earliest",
		"partition":               "0",
		"auto.commit.enable":      "true",
		"auto.commit.interval.ms": "1000",
	}, message3)

	waitGroup.Wait()
}
