package main

import (
	"flag"
)

type Config struct {
	Insenced struct {
		RawEndpoint string
		ProtobufEndpoint string
		EnableProtobuf bool
	}
}


func ParseConfig() (*Config, error) {
	rawEndpoint := flag.String("raw-endpoint", "tcp://*:9424", "ZMQ URI to publish raw frames to.")
	protobufEndpoint := flag.String("protobuf-endpoint", "tcp://*:9425", "ZMQ URI to publish protobufs to (only relevant if -proto is set).")
	enableProtobuf := flag.Bool("enable-protobuf", true, "Enable protobuf output.")
	flag.Parse()
	var cfg Config
	cfg.Insenced.RawEndpoint = *rawEndpoint
	cfg.Insenced.ProtobufEndpoint = *protobufEndpoint
	cfg.Insenced.EnableProtobuf = *enableProtobuf
	return &cfg, nil
}
