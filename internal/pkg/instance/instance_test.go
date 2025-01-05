package instance

import (
	"context"
	"easypwn/internal/data"
	"easypwn/internal/pkg/project"
	"easypwn/internal/pkg/user"
	"fmt"
	"io"
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

	exec, err := instance.Execute(context.Background(), "/bin/bash")
	if err != nil {
		t.Fatal("Failed to execute command: ", err)
	}

	_, err = fmt.Fprintf(exec.Writer, "echo 'Hello, World!'\n")
	if err != nil {
		t.Fatal("Failed to write to PTY: ", err)
	}
	_, err = fmt.Fprintf(exec.Writer, "exit\n")
	if err != nil {
		t.Fatal("Failed to write to PTY: ", err)
	}

	output, err := io.ReadAll(exec.Reader)
	if err != nil {
		t.Fatal("Failed to read from PTY: ", err)
	}
	t.Logf("Result: %s", output)

	err = instance.Stop()
	if err != nil {
		t.Fatal("Failed to stop instance: ", err)
	}

	instance.Delete(context.Background(), data.GetDB())
}
