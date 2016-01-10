package gcan

import (
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/tmc/gcan/gcanpb"
	"golang.org/x/net/context"
)

func newInmemStorage() Storage {
	return &inmemStorage{
		topics:      make(map[string][]*gcanpb.MessageSet, 0),
		subscribers: make(map[chan *gcanpb.MessageSet]bool),
	}
}

type inmemStorage struct {
	// protects the following
	mu     sync.Mutex
	topics map[string][]*gcanpb.MessageSet

	subscribers map[chan *gcanpb.MessageSet]bool
}

func (i *inmemStorage) Publish(ctx context.Context, topic string, ms *gcanpb.MessageSet) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	spew.Dump("Publish:", ms)
	if _, existing := i.topics[topic]; !existing {
		i.topics[topic] = []*gcanpb.MessageSet{}
	}
	ms.Offset = int64(len(i.topics[topic]))
	i.topics[topic] = append(i.topics[topic], ms)

	for subscriber, _ := range i.subscribers {
		go func() {
			subscriber <- ms
		}()
	}
	return nil
}

func (i *inmemStorage) Subscribe(ctx context.Context, topic string, partition uint32, offset uint64) chan *gcanpb.MessageSet {
	i.mu.Lock()
	defer i.mu.Unlock()

	messages := make(chan *gcanpb.MessageSet)
	results := make(chan *gcanpb.MessageSet)

	i.subscribers[messages] = true
	go func() {
		for {
			select {
			case <-ctx.Done():
				i.mu.Lock()
				delete(i.subscribers, messages)
				i.mu.Unlock()
			case ms := <-messages:
				spew.Dump("Subscribe:", ms)
				select {
				case <-ctx.Done():
					close(results)
					results = nil
				case results <- ms:
				}

			}
		}
	}()
	return results
}
