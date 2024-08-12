package ports

import "github.com/sandronister/go_broker/pkg/payload"

type ConfigMap map[string]string

type IBroker interface {
	Consume(config ConfigMap, message chan<- payload.Message) error
	Produce(message *payload.Message, flush int) error
}
