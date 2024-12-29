package instance

import (
	"context"
	"easypwn/assets/images"
	"easypwn/internal/utils"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	cli *client.Client
)

func init() {
	host, exists := os.LookupEnv("DOCKER_HOST")
	if !exists {
		log.Println("DOCKER_HOST is not set, running in test mode")
		return
	}

	initDockerDaemon(host)
	initImages()
}

func initDockerDaemon(host string) {
	ctx := context.Background()

	cli, err := newDockerClient(host)
	if err != nil {
		log.Fatal("Failed to create Docker client: ", err)
	}

	_, err = cli.Ping(ctx)
	if err != nil {
		log.Fatal("Docker daemon is not running: ", err)
	}
}

func initImages() {
	files, err := images.Dockerfiles.ReadDir(".")
	if err != nil {
		log.Fatal("Failed to read Dockerfiles: ", err)
	}

	for _, file := range files {
		dockerfile, err := images.Dockerfiles.ReadFile(file.Name())
		if err != nil {
			log.Fatal("Failed to read Dockerfile: ", err)
		}

		dockerfileTar, err := utils.CreateDockerfileTar(file.Name(), dockerfile)
		if err != nil {
			log.Fatal("Failed to create Dockerfile tar: ", err)
		}

		err = buildDockerImage(context.Background(), cli, dockerfileTar, types.ImageBuildOptions{
			Dockerfile: file.Name(),
			Tags:       []string{utils.DockerfileToImageName(file.Name())},
			Remove:     true,
		})
		if err != nil {
			log.Fatal("Failed to build image: ", err)
		}
	}
}
