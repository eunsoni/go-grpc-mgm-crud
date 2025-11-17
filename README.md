# My Go gRPC Service

This project is a gRPC service implemented in Go. It provides a framework for building efficient and scalable microservices using gRPC and Protocol Buffers.

## Project Structure

```
my-go-grpc-service
├── api
│   └── proto
│       └── service.proto       # Defines the gRPC service and message types
├── cmd
│   └── server
│       └── main.go             # Entry point of the application
├── internal
│   ├── server
│   │   └── server.go           # Implementation of the gRPC server
│   └── service
│       └── service.go          # Business logic for the gRPC service
├── pkg
│   └── middleware
│       └── auth.go             # Middleware functions for authentication
├── scripts
│   └── gen-proto.sh            # Script to generate Go code from .proto file
├── configs
│   └── config.yaml             # Configuration settings for the application
├── Dockerfile                   # Instructions for building a Docker image
├── Makefile                     # Build commands and tasks
├── go.mod                       # Module dependencies
├── go.sum                       # Checksums for module dependencies
├── .gitignore                   # Files and directories to ignore by Git
└── README.md                    # Documentation for the project
```

요구사항
- Go 1.20 이상 (권장 1.24)
- protoc (Protocol Buffers 컴파일러)
- protoc-gen-go, protoc-gen-go-grpc (Go 플러그인)

설치 및 실행 (로컬 개발)
```bash
# 프로젝트 루트로 이동
cd /home/eunseon/go_study/gRPC-user-crud/go-grpc-mgm-crud/my-go-grpc-service

# 의존성 정리
go mod tidy

# protoc 플러그인(한 번만)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin:$(go env GOBIN)"

# .proto에서 Go 코드 생성 (scripts/gen-proto.sh 사용 가능)
# scripts/gen-proto.sh이 있으면: