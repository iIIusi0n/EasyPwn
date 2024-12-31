package instance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BuildOutput struct {
	Stream      string      `json:"stream"`
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

func newDockerClient(host string) (*client.Client, error) {
	return client.NewClientWithOpts(
		client.WithHost(host),
		client.WithAPIVersionNegotiation(),
	)
}

func buildDockerImage(ctx context.Context, cli *client.Client, dockerfileTar io.Reader, options types.ImageBuildOptions) error {
	resp, err := cli.ImageBuild(ctx, dockerfileTar, options)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var output BuildOutput
		if err := decoder.Decode(&output); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if output.Error != "" {
			return fmt.Errorf("build error: %s", output.Error)
		}
	}

	return nil
}

func removeDockerImage(ctx context.Context, cli *client.Client, imageID string) error {
	_, err := cli.ImageRemove(ctx, imageID, image.RemoveOptions{
		Force:         true,
		PruneChildren: true,
	})
	return err
}

func getImages(ctx context.Context, cli *client.Client) ([]image.Summary, error) {
	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return nil, err
	}

	return images, nil
}

func getImageNames(ctx context.Context, cli *client.Client) ([]string, error) {
	images, err := getImages(ctx, cli)
	if err != nil {
		return nil, err
	}

	imageNames := make([]string, len(images))
	for i, image := range images {
		imageNames[i] = image.RepoTags[0]
	}
	return imageNames, nil
}

func createContainer(ctx context.Context, cli *client.Client, containerName, imageName, workPath string) (string, error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"/bin/bash"},
	}, &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/work", workPath),
		},
	}, nil, nil, containerName)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
