package http

import "context"

type ConfigInterface interface {
	GetPort() string
	GetAllowIp() string
}

type ServiceInterface interface {
	ValidateRequest(token string, body []byte) error
	WriteRequest(ctx context.Context, token string, body []byte) error
}
