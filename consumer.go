package gcan

import (
	"fmt"

	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Consumer is the client of the gcan Consumer service.
type Consumer struct{}

var _ gcanpb.ConsumerClient = (*Consumer)(nil)

// Receive sets up a stream to recieve messages given a ReceiveRequest.
func (c *Consumer) Receive(ctx context.Context, req *gcanpb.ReceiveRequest, opts ...grpc.CallOption) (gcanpb.Consumer_ReceiveClient, error) {
	return nil, fmt.Errorf("ConsumerClient.Receive: not implemented")
}
