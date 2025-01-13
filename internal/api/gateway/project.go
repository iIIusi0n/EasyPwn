package gateway

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/pkg/project"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetOsListHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	type OsResponse struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	return func(c *gin.Context) {
		var osList []OsResponse
		res, err := projectClient.GetOsList(context.Background(), &pb.GetOsListRequest{})
		if err != nil {
			log.Printf("Failed to get os list: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get os list"})
			return
		}

		for _, os := range res.OsList {
			osList = append(osList, OsResponse{Id: os.Id, Name: os.Name})
		}
		c.JSON(http.StatusOK, osList)
	}
}

func GetPluginListHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	type PluginResponse struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	return func(c *gin.Context) {
		var pluginList []PluginResponse
		res, err := projectClient.GetPluginList(context.Background(), &pb.GetPluginListRequest{})
		if err != nil {
			log.Printf("Failed to get plugin list: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plugin list"})
			return
		}
		for _, plugin := range res.PluginList {
			pluginList = append(pluginList, PluginResponse{Id: plugin.Id, Name: plugin.Name})
		}
		c.JSON(http.StatusOK, pluginList)
	}
}

func GetProjectsHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	type ProjectResponse struct {
		ProjectId  string `json:"project_id"`
		Name       string `json:"name"`
		UserId     string `json:"user_id"`
		FilePath   string `json:"file_path"`
		FileName   string `json:"file_name"`
		OsName     string `json:"os_name"`
		PluginName string `json:"plugin_name"`
		CreatedAt  string `json:"created_at"`
	}

	return func(c *gin.Context) {
		var projects []ProjectResponse
		res, err := projectClient.GetProjects(context.Background(), &pb.GetProjectsRequest{
			UserId: c.GetString("user_id"),
		})
		if err != nil {
			log.Printf("Failed to get project list: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project list"})
			return
		}
		for _, proj := range res.Projects {
			osName, err := project.GetOsNameFromID(proj.OsId)
			if err != nil {
				log.Printf("Failed to get os: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get os"})
				return
			}

			pluginName, err := project.GetPluginNameFromID(proj.PluginId)
			if err != nil {
				log.Printf("Failed to get plugin: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plugin"})
				return
			}

			projects = append(projects, ProjectResponse{
				ProjectId:  proj.ProjectId,
				Name:       proj.Name,
				UserId:     proj.UserId,
				FilePath:   proj.FilePath,
				FileName:   proj.FileName,
				OsName:     osName,
				PluginName: pluginName,
				CreatedAt:  proj.CreatedAt,
			})
		}
		c.JSON(http.StatusOK, projects)
	}
}

func CreateProjectHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	type CreateProjectRequest struct {
		ProjectName string `form:"project_name" binding:"required"`
		OsId        string `form:"os_id" binding:"required"`
		PluginId    string `form:"plugin_id" binding:"required"`
	}

	type CreateProjectResponse struct {
		ProjectId string `json:"project_id"`
	}

	os.MkdirAll("/var/lib/easypwn/projects", 0755)

	return func(c *gin.Context) {
		var req CreateProjectRequest
		if err := c.ShouldBind(&req); err != nil {
			log.Printf("Failed to bind request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			log.Printf("Failed to get file: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}

		projectDir, err := os.MkdirTemp("/var/lib/easypwn/projects", "easypwn-*")
		if err != nil {
			log.Printf("Failed to create temporary directory: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary directory"})
			return
		}

		if err := c.SaveUploadedFile(file, filepath.Join(projectDir, file.Filename)); err != nil {
			log.Printf("Failed to save file: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		projectDir = strings.Replace(projectDir, "/var/lib/easypwn", dockerHostMountPath, 1)

		res, err := projectClient.CreateProject(context.Background(), &pb.CreateProjectRequest{
			Name:     req.ProjectName,
			UserId:   c.GetString("user_id"),
			FilePath: projectDir,
			FileName: file.Filename,
			OsId:     req.OsId,
			PluginId: req.PluginId,
		})
		if err != nil {
			log.Printf("Failed to create project: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
			return
		}

		c.JSON(http.StatusOK, CreateProjectResponse{
			ProjectId: res.ProjectId,
		})
	}
}

func DeleteProjectHandler(projectClient pb.ProjectClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")
		_, err := projectClient.DeleteProject(context.Background(), &pb.DeleteProjectRequest{
			ProjectId: projectId,
		})
		if err != nil {
			log.Printf("Failed to delete project: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
	}
}
