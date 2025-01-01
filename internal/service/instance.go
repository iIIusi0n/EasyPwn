package service

import (
	"context"

	pb "easypwn/internal/api"
	"easypwn/internal/data"
	"easypwn/internal/pkg/instance"
)

type InstanceService struct {
	pb.UnimplementedInstanceServer
}

func NewInstanceService(ctx context.Context) *InstanceService {
	return &InstanceService{}
}

func (s *InstanceService) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*pb.CreateInstanceResponse, error) {
	instance, err := instance.NewInstance(ctx, data.GetDB(), req.ProjectId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateInstanceResponse{InstanceId: instance.ID}, nil
}

func (s *InstanceService) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.GetInstanceResponse, error) {
	instance, err := instance.GetInstance(ctx, data.GetDB(), req.InstanceId)
	if err != nil {
		return nil, err
	}

	return &pb.GetInstanceResponse{
		InstanceId:  instance.ID,
		ProjectId:   instance.ProjectID,
		ContainerId: instance.ContainerID,
	}, nil
}

func (s *InstanceService) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*pb.DeleteInstanceResponse, error) {
	instance, err := instance.GetInstance(ctx, data.GetDB(), req.InstanceId)
	if err != nil {
		return nil, err
	}

	err = instance.Delete(ctx, data.GetDB())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteInstanceResponse{InstanceId: req.InstanceId}, nil
}

func (s *InstanceService) GetInstanceLogs(ctx context.Context, req *pb.GetInstanceLogsRequest) (*pb.GetInstanceLogsResponse, error) {
	db := data.GetDB()

	instance, err := instance.GetInstance(ctx, db, req.InstanceId)
	if err != nil {
		return nil, err
	}

	logs, err := instance.GetLogs(ctx, db, int(req.Limit))
	if err != nil {
		return nil, err
	}

	return &pb.GetInstanceLogsResponse{Logs: logs}, nil
}