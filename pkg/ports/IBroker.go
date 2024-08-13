package ports

import (
	"errors"

	"github.com/sandronister/go_broker/pkg/payload"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

type ConfigMap map[string]string

type IBroker interface {
	Consume(config ConfigMap, message chan<- payload.Message) error
	Produce(message *payload.Message, flush int) error
}
