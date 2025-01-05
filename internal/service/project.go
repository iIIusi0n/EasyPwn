package service

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"easypwn/internal/pkg/project"
	"os"
)

type ProjectService struct {
	pb.UnimplementedProjectServer
}

func NewProjectService(ctx context.Context) *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	db := data.GetDB()

	project, err := project.NewProject(ctx, db, req.Name, req.UserId, req.FilePath, req.FileName, req.Os, req.Plugin)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProjectResponse{ProjectId: project.ID}, nil
}

func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	db := data.GetDB()

	project, err := project.GetProject(ctx, db, req.ProjectId)
	if err != nil {
		return nil, err
	}

	return &pb.GetProjectResponse{
		ProjectId: project.ID,
		Name:      project.Name,
		UserId:    project.UserID,
		FilePath:  project.FilePath,
		Os:        project.OsID,
		Plugin:    project.PluginID,
	}, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	db := data.GetDB()

	project, err := project.GetProject(ctx, db, req.ProjectId)
	if err != nil {
		return nil, err
	}

	err = project.Delete(ctx, db)
	if err != nil {
		return nil, err
	}

	os.RemoveAll(project.FilePath)

	return &pb.DeleteProjectResponse{ProjectId: req.ProjectId}, nil
}
