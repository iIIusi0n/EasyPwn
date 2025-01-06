package gateway

import (
	"context"
	pb "easypwn/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOsListHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := projectClient.GetOsList(context.Background(), &pb.GetOsListRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func GetPluginListHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := projectClient.GetPluginList(context.Background(), &pb.GetPluginListRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
