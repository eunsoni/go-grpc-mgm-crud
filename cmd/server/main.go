package main

import (
	"log"
	"net"

	pb "github.com/eunsoni/go-grpc-mgm-crud/my-go-grpc-service/api/proto"
	"github.com/eunsoni/go-grpc-mgm-crud/my-go-grpc-service/internal/server"

	"google.golang.org/grpc"
)

func main() {
	// Set up a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server implementation
	pb.RegisterMyServiceServer(s, server.NewServer())

	log.Println("Starting gRPC server on port 50051...")
	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
