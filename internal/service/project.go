package service

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"fmt"
)

type ProjectService struct {
	pb.UnimplementedProjectServer
}

func NewProjectService(ctx context.Context) *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	db := data.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var projectID string
	result := tx.QueryRow("INSERT INTO project (name, user_id, file_path, os_id, plugin_id) VALUES (?, ?, ?, ?, ?) RETURNING id", req.Name, req.UserId, req.FilePath, req.Os, req.Plugin)
	err = result.Scan(&projectID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.CreateProjectResponse{ProjectId: projectID}, nil
}

func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	db := data.GetDB()

	var project pb.GetProjectResponse
	err := db.QueryRow("SELECT id, name, user_id, file_path, os_id, plugin_id FROM project WHERE id = ?", req.ProjectId).Scan(&project.ProjectId, &project.Name, &project.UserId, &project.FilePath, &project.Os, &project.Plugin)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	db := data.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec("DELETE FROM project WHERE id = ?", req.ProjectId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("project not found")
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.DeleteProjectResponse{ProjectId: req.ProjectId}, nil
}
