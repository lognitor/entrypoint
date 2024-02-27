package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type Kafka struct {
	writer *kafka.Writer
}

func NewKafka(writer *kafka.Writer) *Kafka {
	return &Kafka{
		writer: writer,
	}
}

func (k *Kafka) WriteMessage(ctx context.Context, key string, message []byte, t time.Time) error {
	return k.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: message,
		Time:  t,
	})
}

func GetDefaultWriter(config ConfigInterface) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(config.GetBrokers()...),
		Topic:                  config.GetTopic(),
		WriteTimeout:           time.Second * 10,
		Async:                  false,
		AllowAutoTopicCreation: true,
	}
}
