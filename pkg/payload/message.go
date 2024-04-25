package payload

import "time"

type Message struct {
	TopicPartition string
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	Headers        []Header
}
