package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log/slog"
	"os"
)

type Server struct {
	config  ConfigInterface
	handler *Handler
	logger  *slog.Logger
}

func NewServer(config ConfigInterface, srv ServiceInterface) (*Server, error) {
	h, err := NewHandler(srv)
	if err != nil {
		return nil, err
	}

	return &Server{
		config:  config,
		handler: h,
		logger:  slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}, nil
}

func (s *Server) GetAddress() string {
	return fmt.Sprintf("%s:%s", s.config.GetAllowIp(), s.config.GetPort())
}

func (s *Server) Start() error {
	s.logger.Info(fmt.Sprintf("Http server started at %s", s.GetAddress()))
	return fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", s.config.GetAllowIp(), s.config.GetPort()), s.handler.Handle)
}
