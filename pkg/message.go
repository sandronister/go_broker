package pkg

import "time"

type SendMessage struct {
	TopicPartition string
	Value          string
	Headers        string
}

type ReceiptMessage struct {
	TopicPartition string
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	Headers        []string
}
