package gateway

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/pkg/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(userClient pb.UserClient) gin.HandlerFunc {
	ctx := context.Background()

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type LoginResponse struct {
		Token string `json:"token"`
	}

	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		res, err := userClient.AuthLogin(ctx, &pb.AuthLoginRequest{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
			return
		}

		token, err := auth.NewToken(res.UserId, req.Email).Encode()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{
			Token: token,
		})
	}
}

func ConfirmHandler(mailer pb.MailerClient) gin.HandlerFunc {
	ctx := context.Background()

	type ConfirmRequest struct {
		Email string `json:"email"`
	}

	type ConfirmResponse struct {
		Message string `json:"message"`
	}

	return func(c *gin.Context) {
		var req ConfirmRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		_, err := mailer.SendConfirmationEmail(ctx, &pb.SendConfirmationEmailRequest{
			Email: req.Email,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send confirmation email"})
			return
		}

		c.JSON(http.StatusOK, ConfirmResponse{
			Message: "Confirmation email sent",
		})
	}
}

func RegisterHandler(userClient pb.UserClient, mailer pb.MailerClient) gin.HandlerFunc {
	ctx := context.Background()

	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}

	type RegisterResponse struct {
		Token string `json:"token"`
	}

	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		res, err := mailer.GetConfirmationCode(ctx, &pb.GetConfirmationCodeRequest{
			Email: req.Email,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get confirmation code"})
			return
		}

		if res.Code != req.Code {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid confirmation code"})
			return
		}

		createRes, err := userClient.CreateUser(ctx, &pb.CreateUserRequest{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		token, err := auth.NewToken(createRes.UserId, req.Email).Encode()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
			return
		}

		c.JSON(http.StatusOK, RegisterResponse{
			Token: token,
		})
	}
}

func ValidHandler() gin.HandlerFunc {
	type ValidResponse struct {
		UserID string `json:"user_id"`
		Email  string `json:"email"`
	}

	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(string)
		email := c.MustGet("user_email").(string)

		c.JSON(http.StatusOK, ValidResponse{
			UserID: userID,
			Email:  email,
		})
	}
}
