package instance

import (
	"context"
	"os"
	"testing"

	"easypwn/assets/images"
	"easypwn/internal/pkg/util"

	"github.com/docker/docker/api/types"
)

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

	dockerfileTar, err := util.CreateDockerfileTar("Dockerfile.ubuntu-2410.gef", dockerfile)
	if err != nil {
		t.Fatal("Failed to create Dockerfile tar: ", err)
	}

	imageName := "easypwn/ubuntu-2410:gef"
	err = buildDockerImage(ctx, cli, dockerfileTar, types.ImageBuildOptions{
		Dockerfile: "Dockerfile.ubuntu-2410.gef",
		Tags:       []string{imageName},
		Remove:     true,
	})
	if err != nil {
		t.Fatal("Failed to build image: ", err)
	}
	t.Log("Image built successfully")
}

func TestGetImages(t *testing.T) {
	ctx := context.Background()
	cli, err := newDockerClient("unix:///var/run/docker.sock")
	if err != nil {
		t.Fatal("Failed to create Docker client: ", err)
	}

	images, err := getImages(ctx, cli)
	if err != nil {
		t.Fatal("Failed to get images: ", err)
	}

	for _, image := range images {
		t.Logf("Image: %s", image.ID)
	}
}

func TestGetImageNames(t *testing.T) {
	ctx := context.Background()
	cli, err := newDockerClient("unix:///var/run/docker.sock")
	if err != nil {
		t.Fatal("Failed to create Docker client: ", err)
	}

	imageNames, err := getImageNames(ctx, cli)
	if err != nil {
		t.Fatal("Failed to get image names: ", err)
	}

	t.Logf("Image names: %v", imageNames)
}

func TestCreateContainer(t *testing.T) {
	ctx := context.Background()
	cli, err := newDockerClient("unix:///var/run/docker.sock")
	if err != nil {
		t.Fatal("Failed to create Docker client: ", err)
	}

	tempDir, err := os.MkdirTemp("", "easypwn-test-*")
	if err != nil {
		t.Fatal("Failed to create temp dir: ", err)
	}
	defer os.RemoveAll(tempDir)

	containerID, err := createContainer(ctx, cli, "easypwn-test-container", "easypwn/ubuntu-2410:gef", tempDir, true)
	if err != nil {
		t.Fatal("Failed to create container: ", err)
	}

	err = startContainer(ctx, cli, containerID)
	if err != nil {
		t.Fatal("Failed to start container: ", err)
	}

	stopContainer(ctx, cli, containerID)
	removeContainer(ctx, cli, containerID)
}
