package configs

import (
	"errors"
	"os"
	"strings"
)

type Kafka struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
}

func NewKafka() (*Kafka, error) {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		return nil, errors.New("KAFKA_BROKERS is not set")
	}

	b := strings.Split(brokers, ",")
	if len(b) == 0 {
		return nil, errors.New("KAFKA_BROKERS is not set 2")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		return nil, errors.New("KAFKA_TOPIC is not set")
	}

	return &Kafka{
		Brokers: b,
		Topic:   topic,
	}, nil
}

func (k *Kafka) GetBrokers() []string {
	return k.Brokers
}

func (k *Kafka) GetTopic() string {
	return k.Topic
}
