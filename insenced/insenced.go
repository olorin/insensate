package main

import (
	"flag"
	"github.com/anchor/picolog"
	"github.com/fractalcat/emogo"
	zmq "github.com/pebbe/zmq4"
	"os"
)

var Logger *picolog.Logger

func readFrames(e *emogo.EmokitContext, out chan *emogo.EmokitFrame) {
	for {
		f, err := e.GetFrame()
		if err != nil {
			Logger.Errorf("Error reading frame: %v", err)
			continue
		}
		out <- f
	}
}

func main() {
	listenRaw := flag.String("listen-raw", "tcp://*:9424", "ZMQ URI to publish raw frames to.")
	listenProto := flag.String("listen-proto", "tcp://*:9425", "ZMQ URI to publish protobufs to.")
	flag.Parse()
	_ = *listenProto
	Logger = picolog.NewLogger(picolog.LogDebug, "insenced", os.Stdout)
	eeg, err := emogo.NewEmokitContext(emogo.ConsumerHeadset)
	defer eeg.Shutdown()
	if err != nil {
		Logger.Fatalf("Could not initialize emokit context: %v", err)
	}
	Logger.Debugf("EEG initialized.")
	Logger.Debugf("Detected %d EEG devices connected.", eeg.Count())
	rawSock, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		Logger.Fatalf("Could not create ZMQ socket: %v", err)
	}
	err = rawSock.Bind(*listenRaw)
	if err != nil {
		Logger.Fatalf("Could not bind to %s: %v", listenRaw, err)
	}
	frameChan := make(chan *emogo.EmokitFrame)
	go readFrames(eeg, frameChan)
	var f *emogo.EmokitFrame
	for {
		f  = <-frameChan
		Logger.Debugf("Got frame. Sending.")
		_, err := rawSock.SendBytes(f.Raw(), 0)
		if err != nil {
			Logger.Errorf("Error sending raw frame: %v", err)
		}
	}
}
