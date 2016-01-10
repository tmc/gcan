package gcan

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
)

// ProducerServer implements the server portsion of the gcan Producer service.
type ProducerServer struct{}

var _ gcanpb.ProducerServer = (*ProducerServer)(nil)

// Send responds to Send messages from a Producer service client.
func (p *ProducerServer) Send(ctx context.Context, req *gcanpb.SendRequest) (*gcanpb.SendResponse, error) {
	return nil, fmt.Errorf("ProducerServer.Send: not implemented")
}

// SendStream responds to a stream of Send messages from a Producer service client.
func (p *ProducerServer) SendStream(srv gcanpb.Producer_SendStreamServer) error {
	log.Println("got new stream")
	for {
		msg, err := srv.Recv()
		if err != nil {
			log.Error(err)
			return err
		}
		log.Println("got new message", msg)

		if err := msg.MessageSet.CheckIntegrity(); err != nil {
			log.Error(err)
			return err
		}

		if err := srv.Send(&gcanpb.SendResponse{}); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
