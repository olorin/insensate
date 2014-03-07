package main

import (
	"log"
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
	cfg, err := ParseConfig()
	if err != nil {
		log.Fatalf("Could not start insenced: %v", err)
	}
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
	err = rawSock.Bind(cfg.Insenced.RawEndpoint)
	if err != nil {
		Logger.Fatalf("Could not bind to %s: %v", cfg.Insenced.RawEndpoint, err)
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
