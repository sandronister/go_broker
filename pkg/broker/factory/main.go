package factory

import (
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

func NewBroker(kind string, host string, port int) types.IBroker {
	switch kind {
	case KAFKA:
		return brokerkafka.NewBroker(host, port)
	case REDIS:
		return brokerredis.NewBroker(host, port)
	}
	return nil
}

func GetBroker() types.IBroker {
	if brokerConnector == nil {
		port, _ := strconv.Atoi(os.Getenv("BROKER_PORT"))

		brokerConnector = NewBroker(os.Getenv("BROKER_KIND"), os.Getenv("BROKER_HOST"), port)
	}
	return brokerConnector
}
