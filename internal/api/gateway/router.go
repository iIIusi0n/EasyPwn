package gateway

import (
	"easypwn/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.AuthLogin)
		auth.POST("/register", handler.AuthRegister)
	}

	user := r.Group("/user")
	{
		user.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
		})
	}

	return r
}
