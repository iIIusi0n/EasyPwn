package stream

import (
	"github.com/gin-gonic/gin"

	pb "easypwn/internal/api"
	authjwt "easypwn/internal/pkg/auth"
)

type RouterClients struct {
	ProjectClient  pb.ProjectClient
	InstanceClient pb.InstanceClient
}

func NewRouter(clients RouterClients) *gin.Engine {
	r := gin.Default()
	r.Use(authjwt.AuthMiddleware())
	r.Use(InstanceAuthMiddleware(clients.ProjectClient, clients.InstanceClient))

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
