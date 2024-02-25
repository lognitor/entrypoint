package kafka

import (
	"github.com/segmentio/kafka-go"
	"time"
)

func GetDefaultWriter(config ConfigInterface) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(config.GetBrokers()...),
		Topic:                  config.GetTopic(),
		WriteTimeout:           time.Second * 10,
		Async:                  false,
		AllowAutoTopicCreation: true,
	}
}
