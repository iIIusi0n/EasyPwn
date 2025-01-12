package stream

import (
	"context"
	pb "easypwn/internal/api"
	"fmt"
	"net/http"

	authjwt "easypwn/internal/pkg/auth"

	"github.com/gin-gonic/gin"
)

func InstanceAuthMiddleware(projectClient pb.ProjectClient, instanceClient pb.InstanceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		if c.Param("id") == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Instance ID is required"})
			c.Abort()
			return
		}

		instanceInfo, err := instanceClient.GetInstance(ctx, &pb.GetInstanceRequest{
			InstanceId: c.Param("id"),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get instance"})
			c.Abort()
			return
		}

		projectInfo, err := projectClient.GetProject(ctx, &pb.GetProjectRequest{
			ProjectId: instanceInfo.ProjectId,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project"})
			c.Abort()
			return
		}

		token := c.Query("token")
		if token != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := authjwt.Decode(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if user.UserID != projectInfo.UserId {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("full_path", fmt.Sprintf("/work/%s", projectInfo.FileName))
		c.Next()
	}
}
