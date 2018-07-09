package main

import (
	"net"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"github.com/tmc/gcan"
	"github.com/tmc/gcan/gcanpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type StartOptions struct {
	ListenAddr string `short:"l" long:"listen" description:"Listen address" default:":4226"`
}

var startOptions StartOptions

func init() {
	if _, err := optionsParser.AddCommand("start", "starts the gcand daemon", "", &startOptions); err != nil {
		log.Fatal(err)
	}
}

func (c *StartOptions) Execute(args []string) error {
	lis, err := net.Listen("tcp", c.ListenAddr)
	if err != nil {
		return err
	}

	server, err := gcan.NewServer()
	if err != nil {
		return err
	}

	grpclog.SetLogger(log.New())
	grpcServer := grpc.NewServer()
	gcanpb.RegisterProducerServer(grpcServer, server)
	gcanpb.RegisterConsumerServer(grpcServer, server)
	log.Println("listening on", c.ListenAddr)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c
		grpcServer.Stop()
	}()
	return grpcServer.Serve(lis)
}
