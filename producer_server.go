package gcan

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
)

// ProducerServer implements the server portsion of the gcan Producer service.
type ProducerServer struct {
	Storage Storage
}

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
		log.Println("got new message for", msg.Topic)
		if err := msg.MessageSet.CheckIntegrity(); err != nil {
			log.Error(err)
			return err
		}
		log.Println("compressed?", msg.MessageSet.IsCompressedSet())
		if err := p.Storage.Publish(context.Background(), msg.Topic, msg.MessageSet); err != nil {
			log.Error(err)
			return err
		}
		if msg.MessageSet.IsCompressedSet() {
			msg.MessageSet, err = msg.MessageSet.DecompressSet()
			if err != nil {
				log.Error(err)
				return err
			}
			//re-run integrity checks
			if err := msg.MessageSet.CheckIntegrity(); err != nil {
				log.Error(err)
				return err
			}
		}
		log.Println("messsage passed integrity checks:", msg)

		if err := srv.Send(&gcanpb.SendResponse{}); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
