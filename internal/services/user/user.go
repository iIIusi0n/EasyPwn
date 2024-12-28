package api

import (
	"context"

	pb "easypwn/internal/api"
)

type Service struct {
	pb.UnimplementedUserServer
}

func NewService(ctx context.Context) *Service {
	return &Service{}
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{
		UserId: "user-123",
	}, nil
}
