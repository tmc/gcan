package gcan

import "github.com/tmc/gcan/gcanpb"

// Server implements the serving side of the high-level client APIs
type Server struct {
	*ProducerServer
	*ConsumerServer
}

// NewServer prepares a new Server instance.
func NewServer() (*Server, error) {
	storage := newInmemStorage()
	return &Server{
		ProducerServer: &ProducerServer{
			Storage: storage,
		},
		ConsumerServer: &ConsumerServer{
			Storage: storage,
		},
	}, nil
}

var _ gcanpb.Server = (*Server)(nil)
