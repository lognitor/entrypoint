package entrypoint

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
	"time"
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

func (s *Server) WriteLogSync(_ context.Context, in *PayloadRequest) (*PayloadReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.validateAndWrite(ctx, in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &PayloadReply{
		Message: "ok",
	}, nil
}

func (s *Server) WriteLogAsync(_ context.Context, in *PayloadRequest) (*PayloadReply, error) {
	go func(srv *Server) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := s.validateAndWrite(ctx, in); err != nil {
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

func (s *Server) validateAndWrite(ctx context.Context, payload *PayloadRequest) error {
	body := map[string]any{
		"level":   payload.Level,
		"prefix":  payload.Prefix,
		"message": payload.Message,
		"trace":   payload.Trace,
		"source":  payload.Source,
		"ip":      payload.IP,
		"time":    payload.Time, //TODO: correct time parse
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err = s.srv.ValidateRequest(payload.Token, b); err != nil {
		return err
	}

	if err = s.srv.WriteRequest(ctx, payload.Token, b); err != nil {
		return err
	}

	return nil
}
