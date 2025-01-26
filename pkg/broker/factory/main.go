package factory

import (
	"errors"
	"os"
	"strconv"

	brokerkafka "github.com/sandronister/go_broker/pkg/broker/broker_kafka"
	brokerredis "github.com/sandronister/go_broker/pkg/broker/broker_redis"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

const (
	KAFKA = "kafka"
	REDIS = "redis"
)

var brokerConnector types.IBroker

func validEnv() error {

	if os.Getenv("BROKER_KIND") == "" {
		return errors.New("BROKER_KIND is not set")
	}

	if os.Getenv("BROKER_HOST") == "" {
		return errors.New("BROKER_HOST is not set")
	}

	if os.Getenv("BROKER_PORT") == "" {
		return errors.New("BROKER_PORT is not set")
	}

	return nil
}

func NewBroker(kind string, host string, port int) types.IBroker {
	switch kind {
	case KAFKA:
		return brokerkafka.NewBroker(host, port)
	case REDIS:
		return brokerredis.NewBroker(host, port)
	}
	return nil
}

func GetBroker() (types.IBroker, error) {

	if err := validEnv(); err != nil {
		return nil, err
	}

	if brokerConnector == nil {
		port, _ := strconv.Atoi(os.Getenv("BROKER_PORT"))

		brokerConnector = NewBroker(os.Getenv("BROKER_KIND"), os.Getenv("BROKER_HOST"), port)
	}

	return brokerConnector, nil
}
