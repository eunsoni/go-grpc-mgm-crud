package server

import (
	"context"
	"log"
	"net"

	pb "github.com/eunsoni/go-grpc-mgm-crud/my-go-grpc-service/api/proto"
	svc "github.com/eunsoni/go-grpc-mgm-crud/my-go-grpc-service/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedMyServiceServer
	s *svc.Service
}

func NewServer() *Server {
	return &Server{s: svc.NewService()}
}

func (s *Server) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(grpcServer, s)

	log.Printf("gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	// Implement graceful shutdown logic if needed
}

// Delegate RPCs to internal service implementation
func (s *Server) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	return s.s.CreateItem(ctx, req)
}

func (s *Server) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	return s.s.GetItem(ctx, req)
}
