package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/context"

	log "github.com/sirupsen/logrus"
	"github.com/tmc/gcan/gcanpb"
)

type CatOptions struct {
	Topic            string `short:"t" long:"topic" description:"Topic to send to." required:"true"`
	KeySeparator     string `short:"k" long:"keySeparator" description:"Value to separate key from values in lines" default:"	"`
	CompressionCodec string `short:"c" long:"compression" description:"Compression codec to use." default:"NONE"`
}

var catOptions CatOptions

func init() {
	validCodecs := []string{}
	for codec := range gcanpb.MessageCompression_value {
		validCodecs = append(validCodecs, codec)
	}
	sort.Strings(validCodecs)
	desc := fmt.Sprintf("valid compression codecs: %s", strings.Join(validCodecs, ", "))
	if _, err := optionsParser.AddCommand("cat", "cat consumes messages from stdin", desc, &catOptions); err != nil {
		log.Fatal(err)
	}
}

func (c *CatOptions) Execute(args []string) error {
	conn, err := globalOptions.client()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := gcanpb.NewProducerClient(conn)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "gcan.topic", c.Topic)
	stream, err := client.SendStream(ctx)
	if err != nil {
		return err
	}

	keySeparator := []byte(c.KeySeparator)
	compressionCodec := gcanpb.MessageCompression(gcanpb.MessageCompression_value[c.CompressionCodec])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf := scanner.Bytes()
		parts := bytes.SplitN(buf, keySeparator, 2)
		var key, value []byte
		key = parts[0]
		if len(parts) > 1 {
			value = parts[1]
		}

		req := &gcanpb.SendRequest{
			Topic:      c.Topic,
			MessageSet: kvToMessageSet(key, value),
		}
		if compressionCodec != gcanpb.Message_NONE {
			req.MessageSet, err = req.MessageSet.Compress(compressionCodec)
		}
		if err != nil {
			return err
		}
		if err := stream.Send(req); err != nil {
			log.Errorln(err)
			return err
		}
		if _, err := stream.Recv(); err != nil {
			log.Errorln(err)
			return err
		}
	}
	return nil
}

func kvToMessageSet(key, value []byte) *gcanpb.MessageSet {
	ms := &gcanpb.MessageSet{
		Messages: []*gcanpb.Message{
			{
				Key:   key,
				Value: value,
			},
		},
	}
	ms.Messages[0].CRC = ms.Messages[0].ComputeCRC()
	return ms
}
