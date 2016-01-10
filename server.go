package gcan

import "github.com/tmc/gcan/gcanpb"

type Server struct {
	*ProducerServer
	*ConsumerServer
}

func NewServer() (*Server, error) {
	return &Server{
		ProducerServer: &ProducerServer{},
		ConsumerServer: &ConsumerServer{},
	}, nil
}

var _ gcanpb.Server = (*Server)(nil)
