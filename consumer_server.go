package gcan

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
)

// ConsumerServer implements the server portsion of the gcan Consumer service.
type ConsumerServer struct {
	Storage Storage
}

var _ gcanpb.ConsumerServer = (*ConsumerServer)(nil)

// Receive responds to consumer client Receive requests.
func (c *ConsumerServer) Receive(req *gcanpb.ReceiveRequest, srv gcanpb.Consumer_ReceiveServer) error {
	// validate and register subscription

	spew.Dump("started subscribe")
	sub := c.Storage.Subscribe(context.Background(), req.Topic, req.Partition, req.Offset)
	spew.Dump("finished subscribe")

	for ms := range sub {
		srv.Send(&gcanpb.ReceiveResponse{
			Topic:      req.Topic,
			Partition:  req.Partition,
			MessageSet: ms,
		})
	}
	return nil
}
