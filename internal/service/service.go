package service

import (
	"context"

	pb "github.com/eunsoni/go-grpc-mgm-crud/my-go-grpc-service/api/proto"
)

// Service represents the gRPC service implementation.
type Service struct{}

func (s *Service) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	// TODO: 실제 저장/아이디 생성 로직 구현
	id := "generated-id"
	return &pb.CreateItemResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
	}, nil
}

// 예: GetItem 구현
func (s *Service) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	// TODO: 실제 조회 로직 구현
	return &pb.GetItemResponse{
		Id:          req.Id,
		Name:        "example-name",
		Description: "example-description",
	}, nil
}
