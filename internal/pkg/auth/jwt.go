package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	UserID string
	Email  string
	Exp    int64
}

func NewToken(userID, email string) *Token {
	return &Token{
		UserID: userID,
		Email:  email,
		Exp:    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
}

func (t *Token) Encode() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": t.UserID,
		"email":   t.Email,
		"exp":     t.Exp,
	})
	return token.SignedString([]byte(globalSecretKey))
}

func Decode(token string) (*Token, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return []byte(globalSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(globalSecretKey), nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("token expired")
	}

	return &Token{
		UserID: claims["user_id"].(string),
		Email:  claims["email"].(string),
		Exp:    int64(exp),
	}, nil
}

func (t *Token) IsExpired() bool {
	return time.Now().Unix() > t.Exp
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := authHeader[7:]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		claims, err := Decode(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
