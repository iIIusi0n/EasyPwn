package service

import (
	"context"
	pb "easypwn/internal/api"
	"testing"
)

func TestCreateProject(t *testing.T) {
	projectService := NewProjectService(context.Background())

	createProjectResponse, err := projectService.CreateProject(context.Background(), &pb.CreateProjectRequest{
		Name:     "Test Project",
		UserId:   "test-user-id",
		FilePath: "/tmp/test",
	})

	if err != nil {
		t.Errorf("CreateProject() error = %v", err)
		return
	}

	getProjectResponse, err := projectService.GetProject(context.Background(), &pb.GetProjectRequest{
		ProjectId: createProjectResponse.ProjectId,
	})

	if err != nil {
		t.Errorf("GetProject() error = %v", err)
		return
	}

	if getProjectResponse.ProjectId != createProjectResponse.ProjectId {
		t.Errorf("GetProject() response.ProjectId = %v, want %v", getProjectResponse.ProjectId, createProjectResponse.ProjectId)
	}

	if getProjectResponse.Name != "Test Project" {
		t.Errorf("GetProject() response.Name = %v, want %v", getProjectResponse.Name, "Test Project")
	}

	if getProjectResponse.UserId != "test-user-id" {
		t.Errorf("GetProject() response.UserId = %v, want %v", getProjectResponse.UserId, "test-user-id")
	}

	if getProjectResponse.FilePath != "/tmp/test" {
		t.Errorf("GetProject() response.FilePath = %v, want %v", getProjectResponse.FilePath, "/tmp/test")
	}

	_, err = projectService.DeleteProject(context.Background(), &pb.DeleteProjectRequest{
		ProjectId: createProjectResponse.ProjectId,
	})

	if err != nil {
		t.Errorf("DeleteProject() error = %v", err)
		return
	}
}

func TestGetNonExistentProject(t *testing.T) {
	projectService := NewProjectService(context.Background())

	_, err := projectService.GetProject(context.Background(), &pb.GetProjectRequest{
		ProjectId: "non-existent-id",
	})

	if err == nil {
		t.Error("GetProject() expected error for non-existent project, got nil")
	}
}

func TestDeleteNonExistentProject(t *testing.T) {
	projectService := NewProjectService(context.Background())

	_, err := projectService.DeleteProject(context.Background(), &pb.DeleteProjectRequest{
		ProjectId: "non-existent-id",
	})

	if err == nil {
		t.Error("DeleteProject() expected error for non-existent project, got nil")
	}
}
