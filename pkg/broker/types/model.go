package types

import (
	"errors"
	"time"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

type Header struct {
	Key   string
	Value string
}

type Message struct {
	Topic     string    `json:"topic"`
	Value     []byte    `json:"value"`
	Key       []byte    `json:"key"`
	Timestamp time.Time `json:"timestamp"`
	Headers   []Header  `json:"headers"`
	GroupID   string    `json:"group_id"`
}

type ConfigBroker struct {
	Topic                []string
	GroupName            string
	ConsumerName         string
	AutoOffsetReset      string
	EnableAutoCommit     bool
	AutoCommitIntervalMS string
	Partition            int
}

type IBroker interface {
	Consumer(conf *ConfigBroker, message chan<- Message) error
	Producer(message *Message) error
}
