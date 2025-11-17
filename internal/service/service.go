package service

import (
	"context"
	"sync"

	pb "github.com/eunsoni/go-grpc-mgm-crud/my-go-grpc-service/api/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service represents the gRPC service implementation.
type Service struct {
	mu    sync.RWMutex
	store map[string]*pb.CreateItemResponse
}

func NewService() *Service {
	return &Service{
		store: make(map[string]*pb.CreateItemResponse),
	}
}

func (s *Service) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	id := uuid.NewString()
	resp := &pb.CreateItemResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	s.mu.Lock()
	s.store[id] = resp
	s.mu.Unlock()

	return resp, nil
}

func (s *Service) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	s.mu.RLock()
	item, ok := s.store[req.Id]
	s.mu.RUnlock()
	if !ok {
		return nil, status.Errorf(codes.NotFound, "item %s not found", req.Id)
	}

	return &pb.GetItemResponse{
		Id:          item.Id,
		Name:        item.Name,
		Description: item.Description,
	}, nil
}
