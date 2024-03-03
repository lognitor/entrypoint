package grpc

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
	"os"
)

type Server struct {
	UnimplementedEntrypointServer
	config ConfigInterface
	srv    ServiceInterface
	logger *slog.Logger
}

func NewServer(config ConfigInterface, srv ServiceInterface) *Server {
	return &Server{
		config: config,
		srv:    srv,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (s *Server) WriteLogSync(ctx context.Context, in *PayloadRequest) (*PayloadReply, error) {
	if err := s.validateAndWrite(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &PayloadReply{
		Message: "ok",
	}, nil
}

func (s *Server) WriteLogAsync(ctx context.Context, in *PayloadRequest) (*PayloadReply, error) {
	go func(srv *Server) {
		if err := s.validateAndWrite(in); err != nil {
			srv.logger.Error(fmt.Sprintf("Handler error: %s", err.Error()))
		}
	}(s)

	return &PayloadReply{
		Message: "ok",
	}, nil
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%s", s.config.GetAllowIp(), s.config.GetPort())
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterEntrypointServer(grpcServer, s)

	s.logger.Info(fmt.Sprintf("Grpc server started at %s", addr))
	return grpcServer.Serve(lis)
}

func (h *Server) validateAndWrite(payload *PayloadRequest) error {
	body := map[string]any{
		"level":      payload.Level,
		"prefix":     payload.Prefix,
		"message":    payload.Message,
		"lang":       payload.Lang,
		"user_agent": payload.UserAgent,
		"trace":      payload.Trace,
		"ip":         payload.IP,
		"time":       payload.Time, //TODO: correct time parse
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err := h.srv.ValidateRequest(payload.Token, b); err != nil {
		return err
	}

	if err := h.srv.WriteRequest(context.Background(), payload.Token, b); err != nil {
		return err
	}

	return nil
}
