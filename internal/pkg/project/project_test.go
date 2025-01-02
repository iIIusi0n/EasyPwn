package project

import (
	"context"
	"easypwn/internal/data"
	"easypwn/internal/pkg/user"
	"os"
	"testing"
)

func TestProject(t *testing.T) {
	u, err := user.NewUser(context.Background(), data.GetDB(), "test-email", "test-password")
	if err != nil {
		t.Fatal("Failed to create user: ", err)
	}
	defer u.Delete(context.Background(), data.GetDB())

	ubuntu2410, err := GetOsIDFromName("ubuntu-2410")
	if err != nil {
		t.Fatal("Failed to get ubuntu-2410 os ID: ", err)
	}

	gef, err := GetPluginIDFromName("gef")
	if err != nil {
		t.Fatal("Failed to get gef plugin ID: ", err)
	}

	tempDir, err := os.MkdirTemp("", "easypwn-project-test")
	if err != nil {
		t.Fatal("Failed to create temp directory: ", err)
	}
	defer os.RemoveAll(tempDir)

	project, err := NewProject(context.Background(), data.GetDB(), "test-project", u.ID, tempDir, ubuntu2410, gef)
	if err != nil {
		t.Fatal("Failed to create project: ", err)
	}
	defer project.Delete(context.Background(), data.GetDB())

	t.Logf("Project created: %+v", project)
}
