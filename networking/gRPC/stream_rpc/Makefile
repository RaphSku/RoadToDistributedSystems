MIN_GO_VERSION=1.20
MIN_PROTOC_VERSION=3.12.4

.PHONY: default
default: help

.PHONY: all
## Run all
all: main

.PHONY: install_dependencies
## We will need the protobuf-compiler for compiling our .proto files
install_dependencies:
	sudo apt-get update
	sudo apt install -y protobuf-compiler

.PHONY: preflight
## Check that Go and protoc are installed
preflight:
	@if ! test -x "$(shell which go)"; then \
		echo "Go is not installed. Please install it. Minimum Version required: $(MIN_GO_VERSION)"; \
		exit 1; \
	elif [ "$$(($(shell go version | cut -d ' ' -f2 | cut -d '.' -f1 | grep -E -o "[0-9]+") >= $(shell echo $(MIN_GO_VERSION) | cut -d '.' -f1)))" -eq 0 ] || \
		[ "$$(($(shell go version | cut -d ' ' -f2 | cut -d '.' -f2) >= $(shell echo $(MIN_GO_VERSION) | cut -d '.' -f2)))" -eq 0 ] || \
		[ "$$(($(shell go version | cut -d ' ' -f2 | cut -d '.' -f3 | grep -E -o "^[0-9]+") >= $(shell echo $(MIN_GO_VERSION) | cut -d '.' -f3)))" -eq 0 ]; then \
		echo "go version $(MIN_GO_VERSION) or greater is required!"; \
		exit 1; \
	fi
	@if ! test -x "$(shell which go)"; then \
		echo "protoc is not installed. Please install it. Minimum Version required: $(MIN_PROTOC_VERSION)"; \
		exit 1; \
	elif [ "$$(($(shell go version | cut -d ' ' -f2 | cut -d '.' -f1 | grep -E -o "[0-9]+") >= $(shell echo $(MIN_PROTOC_VERSION) | cut -d '.' -f1)))" -eq 0 ] || \
		[ "$$(($(shell go version | cut -d ' ' -f2 | cut -d '.' -f2) >= $(shell echo $(MIN_PROTOC_VERSION) | cut -d '.' -f2)))" -eq 0 ] || \
		[ "$$(($(shell go version | cut -d ' ' -f2 | cut -d '.' -f3 | grep -E -o "^[0-9]+") >= $(shell echo $(MIN_PROTOC_VERSION) | cut -d '.' -f3)))" -eq 0 ]; then \
		echo "protoc version $(MIN_PROTOC_VERSION) or greater is required!"; \
		exit 1; \
	fi

.PHONY: install_tools
install_tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

GOPATH = $(shell which go)/..
.PHONY: main
## Target used to compile the proto files
main: preflight install_tools
	@export GOPATH=$(GOPATH); protoc --proto_path=src --go_out=./server/itemstore --go_opt=paths=source_relative itemstore.proto
	@export GOPATH=$(GOPATH); protoc --proto_path=src --go-grpc_out=./server/itemstore --go-grpc_opt=paths=source_relative itemstore.proto

	@export GOPATH=$(GOPATH); protoc --proto_path=src --go_out=./client/itemstore --go_opt=paths=source_relative itemstore.proto
	@export GOPATH=$(GOPATH); protoc --proto_path=src --go-grpc_out=./client/itemstore --go-grpc_opt=paths=source_relative itemstore.proto

.PHONY: start_client
## Start client application
start_client:
	@cd client && go run main.go

.PHONY: start_server
## Start server application
start_server:
	@cd server && go run main.go

.PHONY: help
## Print this help screen
help:
	@echo "----------------------------------"
	@echo "Welcome to make! Enjoy the flight."
	@echo "Makefile - make [\033[38;5;154mtarget\033[0m]"
	@echo "----------------------------------"
	@echo
	@echo "Targets:"
	@awk '/^[a-zA-z\-_0-9%:\\]+/ { \
		description = match(descriptionLine, /^## (.*)/); \
		if (description) { \
			target = $$1; \
			description = substr(descriptionLine, RSTART + 3, RLENGTH); \
			gsub("\\\\", "", target); \
			gsub(":+$$", "", target); \
			printf "    \033[38;5;154m%-35s\033[0m %s\n", target, description; \
		} \
	} \
	{ descriptionLine = $$0 }' $(MAKEFILE_LIST)
	@printf "\n"
