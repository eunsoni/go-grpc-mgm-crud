#!/bin/bash

# Generate Go code from the .proto file
protoc -I api/proto/ api/proto/service.proto --go_out=plugins=grpc:internal/service