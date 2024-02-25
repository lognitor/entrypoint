package configs

import "os"

type HttpServer struct {
	Port    string `json:"port"`
	AllowIp string `json:"allow_ip"`
}

func NewHttpServer() *HttpServer {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	allowIp := os.Getenv("HTTP_ALLOW_IP")
	if allowIp == "" {
		allowIp = "0.0.0.0"
	}

	return &HttpServer{
		Port:    port,
		AllowIp: allowIp,
	}
}

func (s *HttpServer) GetPort() string {
	return s.Port
}

func (s *HttpServer) GetAllowIp() string {
	return s.AllowIp
}
