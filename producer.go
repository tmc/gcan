package gcan

import (
	"fmt"

	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Producer implements the client of the gcan Producer service.
type Producer struct{}

var _ gcanpb.ProducerClient = (*Producer)(nil)

// Send sends a single message to the gcan Producer service.
func (p *Producer) Send(ctx context.Context, req *gcanpb.SendRequest, opts ...grpc.CallOption) (*gcanpb.SendResponse, error) {
	return nil, fmt.Errorf("Send: not implemented")
}

// SendStream sets up a stream to send messages to the gcan Producer service.
func (p *Producer) SendStream(ctx context.Context, opts ...grpc.CallOption) (gcanpb.Producer_SendStreamClient, error) {
	return nil, fmt.Errorf("SendStream: not implemented")
}
