package service

import (
	"context"
	pb "easypwn/internal/api"
	"testing"
)

func TestInstanceService(t *testing.T) {
	instanceService := NewInstanceService(context.Background())

	createInstanceResponse, err := instanceService.CreateInstance(context.Background(), &pb.CreateInstanceRequest{
		ProjectId: "test-project-id",
	})
	if err != nil {
		t.Errorf("CreateInstance() error = %v", err)
		return
	}

	getInstanceResponse, err := instanceService.GetInstance(context.Background(), &pb.GetInstanceRequest{
		InstanceId: createInstanceResponse.InstanceId,
	})
	if err != nil {
		t.Errorf("GetInstance() error = %v", err)
		return
	}

	t.Logf("Instance created: %+v", getInstanceResponse)

	instanceService.DeleteInstance(context.Background(), &pb.DeleteInstanceRequest{
		InstanceId: createInstanceResponse.InstanceId,
	})
}
