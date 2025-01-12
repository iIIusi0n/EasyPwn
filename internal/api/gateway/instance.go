package gateway

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"easypwn/internal/pkg/instance"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInstancesHandler(instanceClient pb.InstanceClient) gin.HandlerFunc {
	ctx := context.Background()

	type InstanceResponse struct {
		InstanceId string `json:"instance_id"`
		Status     string `json:"status"`
		Memory     int    `json:"memory"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}

	return func(c *gin.Context) {
		projectId := c.Query("project_id")
		if projectId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
			return
		}

		instances, err := instanceClient.GetInstances(ctx, &pb.GetInstancesRequest{
			ProjectId: projectId,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project"})
			return
		}

		var instanceList []InstanceResponse
		for _, insresp := range instances.Instances {
			ins, err := instance.GetInstance(ctx, data.GetDB(), insresp.InstanceId)
			if err != nil {
				log.Printf("Failed to get instance: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get instance"})
				return
			}

			status, err := ins.GetStatus(ctx)
			if err != nil {
				log.Printf("Failed to get instance status: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get instance status"})
				return
			}
			memory, err := ins.GetMemoryUsage(ctx)
			if err != nil {
				log.Printf("Failed to get instance memory usage: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get instance memory usage"})
				return
			}
			instanceList = append(instanceList, InstanceResponse{
				InstanceId: insresp.InstanceId,
				Status:     status,
				Memory:     memory,
				CreatedAt:  ins.CreatedAt.Format(time.RFC3339),
				UpdatedAt:  ins.UpdatedAt.Format(time.RFC3339),
			})
		}

		c.JSON(http.StatusOK, instanceList)
	}
}

func CreateInstanceHandler(instanceClient pb.InstanceClient) gin.HandlerFunc {
	ctx := context.Background()

	type CreateInstanceResponse struct {
		InstanceId string `json:"instance_id"`
	}

	return func(c *gin.Context) {
		projectId := c.Query("project_id")
		if projectId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
			return
		}

		response, err := instanceClient.CreateInstance(ctx, &pb.CreateInstanceRequest{
			ProjectId: projectId,
		})
		if err != nil {
			log.Printf("Failed to create instance: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create instance"})
			return
		}

		c.JSON(http.StatusOK, CreateInstanceResponse{
			InstanceId: response.InstanceId,
		})
	}
}

func DeleteInstanceHandler(instanceClient pb.InstanceClient) gin.HandlerFunc {
	ctx := context.Background()

	return func(c *gin.Context) {
		instanceId := c.Param("id")
		if instanceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Instance ID is required"})
			return
		}

		_, err := instanceClient.DeleteInstance(ctx, &pb.DeleteInstanceRequest{
			InstanceId: instanceId,
		})
		if err != nil {
			log.Printf("Failed to delete instance: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete instance"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Instance deleted successfully"})
	}
}

func ActionInstanceHandler(instanceClient pb.InstanceClient) gin.HandlerFunc {
	ctx := context.Background()

	return func(c *gin.Context) {
		instanceId := c.Param("id")
		if instanceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Instance ID is required"})
			return
		}

		action := c.Query("action")
		if action == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Action is required"})
			return
		}

		switch action {
		case "start":
			_, err := instanceClient.StartInstance(ctx, &pb.StartInstanceRequest{InstanceId: instanceId})
			if err != nil {
				log.Printf("Failed to start instance: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start instance"})
				return
			}
		case "stop":
			_, err := instanceClient.StopInstance(ctx, &pb.StopInstanceRequest{InstanceId: instanceId})
			if err != nil {
				log.Printf("Failed to stop instance: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop instance"})
				return
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Instance action performed successfully"})
	}
}
