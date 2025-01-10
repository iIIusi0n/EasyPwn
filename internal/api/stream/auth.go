package stream

import (
	"context"
	pb "easypwn/internal/api"
	"fmt"
	"net/http"

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

		if projectInfo.UserId != c.MustGet("user_id").(string) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("full_path", fmt.Sprintf("/work/%s", projectInfo.FileName))
		c.Next()
	}
}
