package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/tmc/gcan/gcanpb"
)

type TailOptions struct {
	Topic     string `short:"t" long:"topic" description:"Topic to read from." required:"true"`
	Partition uint32 `short:"p" long:"partiition" description:"Topic partition to consume." default:"0"`
	Offset    uint64 `short:"o" long:"offset" description:"Topic offset to resume from." default:"0"`
}

var tailOptions TailOptions

func init() {
	if _, err := optionsParser.AddCommand("tail", "tail consumes messages from gcan and produces them on stdout", "", &tailOptions); err != nil {
		log.Fatal(err)
	}
}

func (c *TailOptions) Execute(args []string) error {
	conn, err := globalOptions.client()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := gcanpb.NewConsumerClient(conn)

	ctx := context.Background()
	stream, err := client.Receive(ctx, &gcanpb.ReceiveRequest{
		Topic:     c.Topic,
		Offset:    c.Offset,
		Partition: c.Partition,
	})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			return err
		}
		ms := resp.MessageSet
		if err := ms.CheckIntegrity(); err != nil {
			log.Error(err)
			continue
		}
		if ms.IsCompressedSet() {
			ms, err = ms.DecompressSet()
			if err != nil {
				log.Error(err)
				continue
			}
		}

		for _, m := range ms.Messages {
			if m.Key != nil {
				os.Stdout.Write(m.Key)
			}
			if m.Key != nil && m.Value != nil {
				fmt.Print("	")
			}
			if m.Value != nil {
				os.Stdout.Write(m.Value)
			}
			fmt.Println()
		}
	}
}
