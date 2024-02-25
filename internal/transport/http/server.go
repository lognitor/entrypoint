package http

type Server struct {
	config ConfigInterface
}

func NewServer(config ConfigInterface) (*Server, error) {
	return &Server{
		config: config,
	}, nil
}

func (s *Server) Start() error {
	return nil
}
