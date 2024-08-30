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
	Value     []byte
	Key       []byte
	Timestamp time.Time
	Headers   []Header
	GroupID   string
}

type ConfigBroker struct {
	Topic                string
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
