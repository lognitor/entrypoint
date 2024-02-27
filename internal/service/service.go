package service

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ValidateRequest(key string, body []byte) error {
	//req := new(Request)
	//key := ctx.Request.Header.Peek("TOKEN")
	//
	//if len(key) == 0 {
	//	h.error(ctx, fmt.Errorf("token is empty"))
	//	return
	//}
	//
	//if err := json.Unmarshal(ctx.PostBody(), req); err != nil {
	//	h.error(ctx, fmt.Errorf("json unmarshal error: %s", err.Error()))
	//	return
	//}
	//
	//if req.Message == "" || req.Level == "" || req.Prefix == "" {
	//	h.error(ctx, fmt.Errorf("message is empty"))
	//	return
	//}
	//
	//if req.Time.IsZero() {
	//	req.Time = time.Now()
	//}
	//
	//h.kafka.WriteMessage(ctx, string(key), req, req.Time)

	return nil
}

func (s *Service) WriteRequest(body []byte) error {
	return nil
}
