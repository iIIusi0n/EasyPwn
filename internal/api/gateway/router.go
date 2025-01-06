package gateway

import (
	pb "easypwn/internal/api"
	jwtauth "easypwn/internal/pkg/auth"

	"github.com/gin-gonic/gin"
)

type RouterClients struct {
	Mailer         pb.MailerClient
	ChatbotClient  pb.ChatbotClient
	UserClient     pb.UserClient
	ProjectClient  pb.ProjectClient
	InstanceClient pb.InstanceClient
}

func NewRouter(clients RouterClients) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", LoginHandler(clients.UserClient))
		auth.POST("/confirm", ConfirmHandler(clients.Mailer))
		auth.POST("/register", RegisterHandler(clients.UserClient, clients.Mailer))
	}

	user := r.Group("/user")
	{
		user.Use(jwtauth.AuthMiddleware())

		user.GET("/valid", ValidHandler())
	}

	project := r.Group("/project")
	{
		project.Use(jwtauth.AuthMiddleware())

		project.GET("/os", GetOsListHandler(clients.ProjectClient))
		project.GET("/plugin", GetPluginListHandler(clients.ProjectClient))

		project.GET("", GetProjectsHandler(clients.ProjectClient))
		project.POST("", CreateProjectHandler(clients.ProjectClient))
		project.DELETE("/:id", DeleteProjectHandler(clients.ProjectClient))
	}

	return r
}
