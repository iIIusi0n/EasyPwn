package service

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"easypwn/internal/pkg/project"
	"os"
	"time"
)

type ProjectService struct {
	pb.UnimplementedProjectServer
}

func NewProjectService(ctx context.Context) *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	db := data.GetDB()

	project, err := project.NewProject(ctx, db, req.Name, req.UserId, req.FilePath, req.FileName, req.OsId, req.PluginId)
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
		OsId:      project.OsID,
		PluginId:  project.PluginID,
	}, nil
}

func (s *ProjectService) GetProjects(ctx context.Context, req *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	db := data.GetDB()

	projects, err := project.GetProjects(ctx, db, req.UserId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetProjectsResponse{}
	for _, project := range projects {
		response.Projects = append(response.Projects, &pb.GetProjectResponse{
			ProjectId: project.ID,
			Name:      project.Name,
			UserId:    project.UserID,
			FilePath:  project.FilePath,
			FileName:  project.FileName,
			OsId:      project.OsID,
			PluginId:  project.PluginID,
			CreatedAt: project.CreatedAt.Format(time.RFC3339),
		})
	}
	return response, nil
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

func (s *ProjectService) GetOsList(ctx context.Context, req *pb.GetOsListRequest) (*pb.GetOsListResponse, error) {
	oss, err := project.GetOsList()
	if err != nil {
		return nil, err
	}

	response := &pb.GetOsListResponse{}
	for _, os := range oss {
		response.OsList = append(response.OsList, &pb.GetOsResponse{Id: os.ID, Name: os.Name})
	}

	return response, nil
}

func (s *ProjectService) GetPluginList(ctx context.Context, req *pb.GetPluginListRequest) (*pb.GetPluginListResponse, error) {
	plugins, err := project.GetPluginList()
	if err != nil {
		return nil, err
	}

	response := &pb.GetPluginListResponse{}
	for _, plugin := range plugins {
		response.PluginList = append(response.PluginList, &pb.GetPluginResponse{Id: plugin.ID, Name: plugin.Name})
	}
	return response, nil
}
