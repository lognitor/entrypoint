package configs

import (
	"os"
	"strconv"
)

type HttpServer struct {
	Port    uint16 `json:"port"`
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

	// Convert string to uint16
	portUint, _ := strconv.ParseUint(port, 10, 16)

	return &HttpServer{
		Port:    uint16(portUint),
		AllowIp: allowIp,
	}
}

func (s *HttpServer) GetPort() string {
	return strconv.FormatInt(int64(s.Port), 10)
}

func (s *HttpServer) GetAllowIp() string {
	return s.AllowIp
}
