package main

import (
	"github.com/fractalcat/emogo"
	"github.com/anchor/picolog"
	zmq "github.com/pebbe/zmq4"
	"os"
	"flag"
	"log/syslog"
)

var Logger *picolog.Logger

func main() {
	listen := flag.String("listen", "tcp://*:9424", "ZMQ URI to listen on.")
	flag.Parse()
	Logger = picolog.NewLogger(syslog.LOG_DEBUG, "insenced", os.Stdout)
	eeg, err := emogo.NewEmokitContext()
	if err != nil {
		Logger.Errorf("Could not initialize emokit context: %v", err)
	}
	sock, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		Logger.Fatalf("Could not create ZMQ socket: %v", err)
	}
	err = sock.Bind(*listen)
	if err != nil {
		Logger.Fatalf("Could not bind to %s: %v", listen, err)
	}
	_ = eeg
}
