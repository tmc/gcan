package gcan

import (
	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
)

type Storage interface {
	Publish(ctx context.Context, topic string, ms *gcanpb.MessageSet) error
	Subscribe(ctx context.Context, topic string, partition uint32, offset uint64) chan *gcanpb.MessageSet
}
