package instance

import (
	"context"
	"easypwn/assets/images"
	"easypwn/internal/pkg/util"
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
		host = "unix:///var/run/docker.sock"
	}

	initDockerDaemon(host)
	initImages()
}

func initDockerDaemon(host string) {
	ctx := context.Background()

	var err error
	cli, err = newDockerClient(host)
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

	imageNames, err := getImageNames(context.Background(), cli)
	if err != nil {
		log.Fatal("Failed to get image names: ", err)
	}

	for _, file := range files {
		found := false
		for _, tag := range imageNames {
			if tag == util.DockerfileToImageName(file.Name()) {
				found = true
				break
			}
		}
		if found {
			continue
		}

		dockerfile, err := images.Dockerfiles.ReadFile(file.Name())
		if err != nil {
			log.Fatal("Failed to read Dockerfile: ", err)
		}

		dockerfileTar, err := util.CreateDockerfileTar(file.Name(), dockerfile)
		if err != nil {
			log.Fatal("Failed to create Dockerfile tar: ", err)
		}

		err = buildDockerImage(context.Background(), cli, dockerfileTar, types.ImageBuildOptions{
			Dockerfile: file.Name(),
			Tags:       []string{util.DockerfileToImageName(file.Name())},
			Remove:     true,
		})
		if err != nil {
			log.Fatal("Failed to build image: ", err)
		}
	}
}
