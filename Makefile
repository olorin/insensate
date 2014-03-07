PROTOBUFS = EpocFrame.pb.go
PROTOC = protoc --go_out=..

install: build check
	go install

build: $(PROTOBUFS)
	go build

deps:
	go get

EpocFrame.pb.go: protobuf/EpocFrame.proto
	cd protobuf ; $(PROTOC) EpocFrame.proto

check: protobuf
	go test

clean:
	rm -f $(PROTOBUFS)

.PHONY : all
.PHONY : clean
