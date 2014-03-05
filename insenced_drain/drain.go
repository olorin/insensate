package main

import (
	zmq "github.com/pebbe/zmq4"
	"flag"
	"os"
	"log"
)

func main() {
	uri := flag.String("insenced", "tcp://127.0.0.1:9424", "ZMQ URI insenced is listening on.")
	flag.Parse()
	sock, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatalf("Could not create ZMQ socket: %v", err)
	}
	err = sock.Connect(*uri)
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", uri, err)
	}
	err = sock.SetSubscribe("")
	if err != nil {
		log.Fatalf("Could not subscribe: %v", uri, err)
	}
	for {
		msg, err := sock.RecvBytes(0)
		if err != nil {
			log.Fatalf("Error receiving EEG frame: %v", err)
		}
		os.Stdout.Write(msg)
	}
}
