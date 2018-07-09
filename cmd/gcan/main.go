package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Be verbose."`
	Server  string `short:"s" long:"bootstrap-server" description:"Server addr to bootstrap with." default:":4226"`
}

var globalOptions Options

var optionsParser = flags.NewNamedParser("gcan", flags.Default)

func init() {
	optionsParser.ShortDescription = "gcan cli"
	optionsParser.AddGroup("Global options", "", &globalOptions)
}

// client attempts to construct a mowgli client from the Server option value
func (o *Options) client() (client *grpc.ClientConn, err error) {
	grpclog.SetLogger(log.New())
	client, err = grpc.Dial(o.Server, grpc.WithTimeout(time.Second), grpc.WithBlock(), grpc.WithInsecure())
	if err == nil {
		return
	}
	if len(o.Verbose) > 0 {
		fmt.Fprintf(os.Stderr, "issue connecting to %v: '%v'\n", o.Server, err)
	}
	if client == nil && err == nil {
		return nil, fmt.Errorf("gcan: could not connect to the provided gcand server (attempted: '%v')", o.Server)
	}
	return client, err
}

func main() {
	if _, err := optionsParser.Parse(); err != nil {
		os.Exit(1)
	}
}
