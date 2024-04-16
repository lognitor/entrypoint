package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(writer *kafka.Writer) *Producer {
	return &Producer{
		writer: writer,
	}
}

func (k *Producer) WriteMessage(ctx context.Context, key string, message []byte, t time.Time) error {
	return k.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: message,
		Time:  t,
	})
}

func GetDefaultProducer(config ConfigInterface) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(config.GetBrokers()...),
		Topic:                  config.GetTopic(),
		WriteTimeout:           time.Second * 10,
		Async:                  false,
		AllowAutoTopicCreation: true,
	}
}
