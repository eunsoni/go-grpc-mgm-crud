# Makefile for my-go-grpc-service

.PHONY: all build run clean proto

all: build

build:
	go build -o my-go-grpc-service ./cmd/server/main.go

run: build
	./my-go-grpc-service

clean:
	go clean
	rm -f my-go-grpc-service

proto:
	protoc --go_out=. --go-grpc_out=. api/proto/service.proto

test:
	go test ./... -v

fmt:
	go fmt ./...