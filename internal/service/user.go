package service

import (
	"context"
	"fmt"

	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"easypwn/internal/pkg/user"
	"easypwn/internal/pkg/util"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{}
}

func (s *UserService) AuthLogin(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	passwordHash := util.HashPassword(req.Password)

	db := data.GetDB()

	user, err := user.GetUserByEmail(ctx, db, req.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != passwordHash {
		return nil, fmt.Errorf("invalid password")
	}

	return &pb.AuthLoginResponse{UserId: user.ID}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	db := data.GetDB()

	user, err := user.NewUser(ctx, db, req.Email, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{UserId: user.ID}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	db := data.GetDB()

	u, err := user.GetUser(ctx, db, req.UserId)
	if err != nil {
		return nil, err
	}

	licenseType, err := u.GetLicense(ctx, db)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		UserId:      u.ID,
		Email:       u.Email,
		Username:    u.Username,
		LicenseType: licenseType,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	db := data.GetDB()

	u, err := user.GetUser(ctx, db, req.UserId)
	if err != nil {
		return nil, err
	}

	if req.Password != "" {
		err = u.UpdatePassword(ctx, db, req.Password)
		if err != nil {
			return nil, err
		}
	}

	if req.LicenseType != "" {
		err = u.UpdateLicense(ctx, db, req.LicenseType)
		if err != nil {
			return nil, err
		}
	}

	return &pb.UpdateUserResponse{UserId: req.UserId}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	db := data.GetDB()

	u, err := user.GetUser(ctx, db, req.UserId)
	if err != nil {
		return nil, err
	}

	err = u.Delete(ctx, db)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{UserId: u.ID}, nil
}
