all: main

# We will need the protobuf-compiler for compiling our .proto files
install_dependencies:
	sudo apt-get update
	sudo apt install -y protobuf-compiler

preflight:
	@which go
	@go version
	@which protoc
	@protoc --version | grep -E '[[:digit:]].[[:digit:]]{1,3}.[[:digit:]]' -o

go_plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

GOPATH = $(shell which go)/bin
main: preflight go_plugins
	@export GOPATH=$(GOPATH); protoc --proto_path=src --go_out=out --go_opt=paths=source_relative itemstore.proto