package pkg

import "time"

type SendMessage struct {
	TopicPartition string
	Value          string
	Headers        []Header
}

type ReceiptMessage struct {
	TopicPartition string
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	Headers        []string
}
