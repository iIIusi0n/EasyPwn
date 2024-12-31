package service

import (
	"context"

	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"easypwn/internal/util"
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

	var userId string
	row := db.QueryRow("SELECT id FROM user WHERE email = ? AND password_hash = ?", req.Email, passwordHash)
	err := row.Scan(&userId)
	if err != nil {
		return nil, err
	}

	return &pb.AuthLoginResponse{
		UserId: userId,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	passwordHash := util.HashPassword(req.Password)

	db := data.GetDB()

	var userId string
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result := tx.QueryRow("INSERT INTO user (email, username, password_hash) VALUES (?, ?, ?) RETURNING id", req.Email, req.Username, passwordHash)
	err = result.Scan(&userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		INSERT INTO user_license (user_id, license_type_id) 
		SELECT ?, id FROM user_license_type WHERE name = 'free'
	`, userId)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		UserId: userId,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	db := data.GetDB()

	row := db.QueryRow(`
		SELECT u.id, u.email, u.username, ult.name as license_type 
		FROM user u
		LEFT JOIN user_license ul ON u.id = ul.user_id
		LEFT JOIN user_license_type ult ON ul.license_type_id = ult.id
		WHERE u.id = ?`, req.UserId)

	var user pb.GetUserResponse
	err := row.Scan(&user.UserId, &user.Email, &user.Username, &user.LicenseType)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	db := data.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE user SET email = ?, username = ? WHERE id = ?", req.Email, req.Username, req.UserId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		UPDATE user_license 
		SET license_type_id = (SELECT id FROM user_license_type WHERE name = ?)
		WHERE user_id = ?
	`, req.LicenseType, req.UserId)
	if err != nil {
		return nil, err
	}

	if req.Password != "" {
		passwordHash := util.HashPassword(req.Password)
		_, err = tx.Exec("UPDATE user SET password_hash = ? WHERE id = ?", passwordHash, req.UserId)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		UserId: req.UserId,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	db := data.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM user WHERE id = ?", req.UserId)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		UserId: req.UserId,
	}, nil
}
