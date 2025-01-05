package gateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register")
	}

	user := r.Group("/user")
	{
		user.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
		})
	}

	return r
}
