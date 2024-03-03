package grpc

import "context"

type configPortInterface interface {
	GetPort() string
}

type allowIpInterface interface {
	GetAllowIp() string
}

type ConfigInterface interface {
	configPortInterface
	allowIpInterface
}

type ServiceInterface interface {
	ValidateRequest(token string, body []byte) error
	WriteRequest(ctx context.Context, token string, body []byte) error
}
