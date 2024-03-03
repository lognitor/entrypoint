package configs

import (
	"os"
	"strconv"
)

type GrpcServer struct {
	Port    uint16 `json:"port"`
	AllowIp string `json:"allow_ip"`
}

func NewGrpcServer() *GrpcServer {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	allowIp := os.Getenv("GRPC_ALLOW_IP")
	if allowIp == "" {
		allowIp = "0.0.0.0"
	}

	// Convert string to uint16
	portUint, _ := strconv.ParseUint(port, 10, 16)

	return &GrpcServer{
		Port:    uint16(portUint),
		AllowIp: allowIp,
	}
}

func (s *GrpcServer) GetPort() string {
	return strconv.FormatInt(int64(s.Port), 10)
}

func (s *GrpcServer) GetAllowIp() string {
	return s.AllowIp
}
