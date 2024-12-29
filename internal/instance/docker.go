package instance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
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
