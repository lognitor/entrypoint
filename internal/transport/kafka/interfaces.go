package kafka

type ConfigInterface interface {
	GetBrokers() []string
	GetTopic() string
}
