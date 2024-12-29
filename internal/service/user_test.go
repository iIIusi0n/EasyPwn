package service

import (
	"context"
	pb "easypwn/internal/api"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userService := NewUserService(context.Background())

	createUserResponse, err := userService.CreateUser(context.Background(), &pb.CreateUserRequest{
		Email:    "test@test.com",
		Username: "test",
		Password: "password123",
	})

	if err != nil {
		t.Errorf("CreateUser() error = %v", err)
		return
	}

	getUserResponse, err := userService.GetUser(context.Background(), &pb.GetUserRequest{
		UserId: createUserResponse.UserId,
	})

	if err != nil {
		t.Errorf("GetUser() error = %v", err)
		return
	}

	if getUserResponse.UserId != createUserResponse.UserId {
		t.Errorf("GetUser() response.UserId = %v, want %v", getUserResponse.UserId, createUserResponse.UserId)
	}

	_, err = userService.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		UserId: createUserResponse.UserId,
	})

	if err != nil {
		t.Errorf("DeleteUser() error = %v", err)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	userService := NewUserService(context.Background())

	createResp, err := userService.CreateUser(context.Background(), &pb.CreateUserRequest{
		Email:    "test@test.com",
		Username: "test",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	_, err = userService.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		UserId:      createResp.UserId,
		Email:       "test@test.com",
		Username:    "test",
		Password:    "",
		LicenseType: "paid",
	})

	if err != nil {
		t.Errorf("UpdateUser() error = %v", err)
		return
	}

	getUserResponse, err := userService.GetUser(context.Background(), &pb.GetUserRequest{
		UserId: createResp.UserId,
	})

	if err != nil {
		t.Errorf("GetUser() error = %v", err)
		return
	}

	if getUserResponse.LicenseType != "paid" {
		t.Errorf("GetUser() response.LicenseType = %v, want %v", getUserResponse.LicenseType, "paid")
	}

	_, err = userService.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		UserId: createResp.UserId,
	})

	if err != nil {
		t.Errorf("DeleteUser() error = %v", err)
		return
	}
}

func TestAuthLogin(t *testing.T) {
	userService := NewUserService(context.Background())

	createResp, err := userService.CreateUser(context.Background(), &pb.CreateUserRequest{
		Email:    "test@test.com",
		Username: "test",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	t.Run("SuccessfulLogin", func(t *testing.T) {
		authLoginResponse, err := userService.AuthLogin(context.Background(), &pb.AuthLoginRequest{
			Email:    "test@test.com",
			Password: "password123",
		})

		if err != nil {
			t.Errorf("AuthLogin() error = %v", err)
			return
		}

		if authLoginResponse.UserId != createResp.UserId {
			t.Errorf("AuthLogin() response.UserId = %v, want %v", authLoginResponse.UserId, createResp.UserId)
		}
	})

	t.Run("NonExistentUser", func(t *testing.T) {
		_, err := userService.AuthLogin(context.Background(), &pb.AuthLoginRequest{
			Email:    "test2@test.com",
			Password: "password123",
		})

		if err == nil {
			t.Error("AuthLogin() expected error for non-existent user, got nil")
		}
	})

	t.Run("WrongPassword", func(t *testing.T) {
		_, err := userService.AuthLogin(context.Background(), &pb.AuthLoginRequest{
			Email:    "test@test.com",
			Password: "wrongpassword",
		})

		if err == nil {
			t.Error("AuthLogin() expected error for wrong password, got nil")
		}
	})

	_, err = userService.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		UserId: createResp.UserId,
	})

	if err != nil {
		t.Errorf("DeleteUser() error = %v", err)
		return
	}
}
