package instance

import (
	"context"
	"easypwn/internal/data"
	"easypwn/internal/pkg/project"
	"easypwn/internal/pkg/user"
	"os"
	"testing"
)

func TestInstance(t *testing.T) {
	u, err := user.NewUser(context.Background(), data.GetDB(), "test-email", "test-password")
	if err != nil {
		t.Fatal("Failed to create user: ", err)
	}
	defer u.Delete(context.Background(), data.GetDB())

	ubuntu2410, err := project.GetOsIDFromName("ubuntu-2410")
	if err != nil {
		t.Fatal("Failed to get ubuntu-2410 os ID: ", err)
	}

	gef, err := project.GetPluginIDFromName("gef")
	if err != nil {
		t.Fatal("Failed to get gef plugin ID: ", err)
	}

	tempDir, err := os.MkdirTemp("", "easypwn-instance-test")
	if err != nil {
		t.Fatal("Failed to create temp directory: ", err)
	}
	defer os.RemoveAll(tempDir)

	project, err := project.NewProject(context.Background(), data.GetDB(), "test-project", u.ID, tempDir, ubuntu2410, gef)
	if err != nil {
		t.Fatal("Failed to create project: ", err)
	}
	defer project.Delete(context.Background(), data.GetDB())

	instance, err := NewInstance(context.Background(), data.GetDB(), project.ID)
	if err != nil {
		t.Fatal("Failed to create instance: ", err)
	}

	t.Logf("Instance created: %+v", instance)

	err = instance.Stop()
	if err != nil {
		t.Fatal("Failed to stop instance: ", err)
	}

	instance.Delete(context.Background(), data.GetDB())
}
