package brokerkafka

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (b *Broker) getHeader(message *types.Message) []kafka.Header {
	var listHeader []kafka.Header

	for _, item := range message.Headers {
		header := kafka.Header{Key: item.Key, Value: []byte(item.Value)}
		listHeader = append(listHeader, header)
	}

	return listHeader
}

func (b *Broker) getConsumer(config *types.ConfigBroker) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(b.GetConfig(config))

	if err != nil {
		return nil, err
	}

	if config.Partition != 0 {

		for _, topic := range config.Topic {
			err := c.Assign([]kafka.TopicPartition{
				{Topic: &topic, Partition: int32(config.Partition)},
			})

			if err != nil {
				return nil, err
			}
		}
	}

	if config.Partition == 0 {
		err := c.SubscribeTopics(config.Topic, nil)

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (b *Broker) ListenToQueue(config *types.ConfigBroker, message chan<- types.Message) error {

	c, err := b.getConsumer(config)

	if err != nil {
		return err
	}

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			return err
		}

		var listHeaders []types.Header

		for _, h := range msg.Headers {
			header := types.Header{Key: h.Key, Value: string(h.Value)}
			listHeaders = append(listHeaders, header)
		}

		message <- types.Message{
			Value:     msg.Value,
			Key:       msg.Key,
			Timestamp: msg.Timestamp,
			Headers:   listHeaders,
		}

	}

}

func (b *Broker) GetConfig(config *types.ConfigBroker) *kafka.ConfigMap {
	if config.GroupName == "" {
		config.GroupName = "default"
	}

	if config.AutoOffsetReset == "" {
		config.AutoOffsetReset = "earliest"
	}

	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", b.host, strconv.Itoa(b.port)),
		"group.id":          config.GroupName,
		"auto.offset.reset": config.AutoOffsetReset,
	}

	if !config.EnableAutoCommit {
		kafkaConfig.SetKey("enable.auto.commit", true)
	}

	if config.AutoCommitIntervalMS != "" {
		kafkaConfig.SetKey("auto.commit.interval.ms", config.AutoCommitIntervalMS)
	}

	return kafkaConfig

}

func (b *Broker) Publish(message *types.Message) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", b.host, strconv.Itoa(b.port))})

	if err != nil {
		return err
	}

	headers := b.getHeader(message)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &message.Topic, Partition: 0},
		Value:          []byte(message.Value),
		Headers:        headers,
	}, nil)

	if err != nil {
		return err
	}

	p.Close()

	return nil
}
