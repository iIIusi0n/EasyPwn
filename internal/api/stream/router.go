package stream

import (
	"github.com/gin-gonic/gin"

	pb "easypwn/internal/api"
	"easypwn/internal/pkg/auth"
)

func NewRouter(projectClient pb.ProjectClient, instanceClient pb.InstanceClient) *gin.Engine {
	r := gin.Default()
	r.Use(auth.AuthMiddleware())
	r.Use(InstanceAuthMiddleware(projectClient, instanceClient))

	stream := r.Group("/stream")
	{
		session := stream.Group("/session")
		{
			session.GET("/debugger/:id", GetDebuggerSessionHandler())

			session.GET("/shell/:id", GetShellSessionHandler())
		}
	}

	return r
}
