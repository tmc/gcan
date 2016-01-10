package gcan

import (
	"fmt"

	"github.com/tmc/gcan/gcanpb"
)

// ConsumerServer implements the server portsion of the gcan Consumer service.
type ConsumerServer struct{}

var _ gcanpb.ConsumerServer = (*ConsumerServer)(nil)

// Receive responds to consumer client Receive requests.
func (c *ConsumerServer) Receive(req *gcanpb.ReceiveRequest, srv gcanpb.Consumer_ReceiveServer) error {
	return fmt.Errorf("ConsumerServer.Receive: not implemented")
}
