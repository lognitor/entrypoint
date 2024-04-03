package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/lognitor/entrypoint/pkg/structs"
	"time"
)

type Service struct {
	w KafkaWriterInterface
}

func NewService(w KafkaWriterInterface) *Service {
	return &Service{
		w: w,
	}
}

func (s *Service) ValidateRequest(token string, body []byte) error {
	if token == "" {
		return errors.New("token is empty")
	}

	if len(body) == 0 {
		return errors.New("body is empty")
	}

	// TODO: Send token for validation to lognitor
	req := &structs.Log{}

	if err := json.Unmarshal(body, req); err != nil {
		return fmt.Errorf("json unmarshal error: %s", err.Error())
	}

	if req.Message == "" || req.Level == "" || req.Prefix == "" {
		return fmt.Errorf("message, level or prefix is empty")
	}

	return nil
}

func (s *Service) WriteRequest(ctx context.Context, token string, body []byte) error {
	req := &structs.Log{}

	if err := json.Unmarshal(body, req); err != nil {
		return fmt.Errorf("json unmarshal error: %s", err.Error())
	}

	if req.Time.IsZero() {
		req.Time = time.Now()
	}

	b, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshal error: %s", err.Error())
	}

	if err = s.w.WriteMessage(ctx, token, b, req.Time); err != nil {
		return fmt.Errorf("kafka write error: %s", err.Error())
	}

	return nil
}
