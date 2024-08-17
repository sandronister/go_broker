package brokerkafka

import "github.com/sandronister/go_broker/pkg/broker/types"

func (b *Broker) Consumer(config *types.ConfigMap, message chan<- types.Message) error {

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
