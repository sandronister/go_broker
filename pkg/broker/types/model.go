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
	TopicPartition string
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	Headers        []Header
}

type ConfigMap map[string]string

type IBroker interface {
	Consumer(conf *ConfigMap, message chan<- Message) error
	Producer(message *Message, flush int) error
}
