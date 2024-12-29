package instance

import (
	"context"
	"easypwn/assets/images"
	"easypwn/internal/utils"
	"testing"

	"github.com/docker/docker/api/types"
)

func TestEmbeddedDockerfiles(t *testing.T) {
	files, err := images.Dockerfiles.ReadDir(".")
	if err != nil {
		t.Fatal("Failed to read Dockerfiles: ", err)
	}

	for _, file := range files {
		_, err := images.Dockerfiles.ReadFile(file.Name())
		if err != nil {
			t.Fatal("Failed to read Dockerfile: ", err)
		}
		t.Logf("Read Dockerfile: %s", file.Name())
	}
}

func TestBuildImage(t *testing.T) {
	dockerfile, err := images.Dockerfiles.ReadFile("Dockerfile.ubuntu-2410.gef")
	if err != nil {
		t.Fatal("Failed to read Dockerfile: ", err)
	}

	ctx := context.Background()
	cli, err := newDockerClient("unix:///var/run/docker.sock")
	if err != nil {
		t.Fatal("Failed to create Docker client: ", err)
	}

	_, err = cli.Ping(ctx)
	if err != nil {
		t.Fatal("Docker daemon is not running: ", err)
	}

	dockerfileTar, err := utils.CreateDockerfileTar("Dockerfile.ubuntu-2410.gef", dockerfile)
	if err != nil {
		t.Fatal("Failed to create Dockerfile tar: ", err)
	}

	imageName := "easypwn/ubuntu-2410/gef"
	err = buildDockerImage(ctx, cli, dockerfileTar, types.ImageBuildOptions{
		Dockerfile: "Dockerfile.ubuntu-2410.gef",
		Tags:       []string{imageName},
		Remove:     true,
	})
	if err != nil {
		t.Fatal("Failed to build image: ", err)
	}
	t.Log("Image built successfully")

	err = removeDockerImage(ctx, cli, imageName)
	if err != nil {
		t.Fatal("Failed to remove image: ", err)
	}
	t.Log("Image removed successfully")
}
