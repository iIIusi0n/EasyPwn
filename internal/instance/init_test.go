package instance

import (
	"bytes"
	"context"
	"easypwn/assets/images"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
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

	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"))
	if err != nil {
		t.Fatal("Failed to create Docker client: ", err)
	}

	ctx := context.Background()
	_, err = cli.Ping(ctx)
	if err != nil {
		t.Fatal("Docker daemon is not running: ", err)
	}

	_, err = cli.ImageBuild(ctx, bytes.NewReader(dockerfile), types.ImageBuildOptions{
		Dockerfile: "Dockerfile.ubuntu-2410.gef",
		Tags:       []string{"easypwn/ubuntu-2410/gef"},
	})
	if err != nil {
		t.Fatal("Failed to build image: ", err)
	}

	t.Log("Image built successfully")

	_, err = cli.ImageRemove(ctx, "easypwn/ubuntu-2410/gef", image.RemoveOptions{
		Force:         true,
		PruneChildren: true,
	})
	if err != nil {
		t.Fatal("Failed to remove image: ", err)
	}

	t.Log("Image removed successfully")
}
