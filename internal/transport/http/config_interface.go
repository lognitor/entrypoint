package http

type ConfigInterface interface {
	GetPort() string
	GetAllowIp() string
}
