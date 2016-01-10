package gcan

import "github.com/tmc/gcan/gcanpb"

// Server implements the serving side of the high-level client APIs
type Server struct {
	*ProducerServer
	*ConsumerServer
}

// NewServer prepares a new Server instance.
func NewServer() (*Server, error) {
	return &Server{
		ProducerServer: &ProducerServer{},
		ConsumerServer: &ConsumerServer{},
	}, nil
}

var _ gcanpb.Server = (*Server)(nil)
