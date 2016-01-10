package main

import (
	"bufio"
	"bytes"
	"os"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/tmc/gcan/gcanpb"
)

type CatOptions struct {
	Topic        string `short:"t" long:"topic" description:"Topic to send to." required:"true"`
	KeySeparator string `short:"k" long:"keySeparator" description:"Value to separate key from values in lines" default:"	"`
}

var catOptions CatOptions

func init() {
	if _, err := optionsParser.AddCommand("cat", "cat consumes messages from stdin", "", &catOptions); err != nil {
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf := scanner.Bytes()
		parts := bytes.SplitN(buf, keySeparator, 2)
		req := &gcanpb.SendRequest{
			Topic: c.Topic,
			Key:   string(parts[0]),
		}
		if len(parts) > 1 {
			req.Value = parts[1]
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
