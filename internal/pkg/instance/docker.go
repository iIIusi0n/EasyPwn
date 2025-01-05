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

type ExecInOut struct {
	ExecID string
	Reader io.Reader
	Writer io.Writer
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

	var imageNames []string
	for _, image := range images {
		imageNames = append(imageNames, image.RepoTags...)
	}
	return imageNames, nil
}

func createContainer(ctx context.Context, cli *client.Client, containerName, imageName, workPath string, autoRemove bool) (string, error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"/bin/bash"},
		Tty:   true,
	}, &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/work", workPath),
		},
		AutoRemove: autoRemove,
	}, nil, nil, containerName)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func startContainer(ctx context.Context, cli *client.Client, containerID string) error {
	return cli.ContainerStart(ctx, containerID, container.StartOptions{})
}

func stopContainer(ctx context.Context, cli *client.Client, containerID string) error {
	timeout := 5
	return cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &timeout})
}

func removeContainer(ctx context.Context, cli *client.Client, containerID string) error {
	return cli.ContainerRemove(ctx, containerID, container.RemoveOptions{
		Force: true,
	})
}

func executeCommand(ctx context.Context, cli *client.Client, containerID string, command ...string) (ExecInOut, error) {
	execID, err := cli.ContainerExecCreate(ctx, containerID, container.ExecOptions{
		Cmd:          command,
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		ConsoleSize:  &[2]uint{24, 80},
		WorkingDir:   "/work",
	})
	if err != nil {
		return ExecInOut{}, fmt.Errorf("failed to create exec: %v", err)
	}

	resp, err := cli.ContainerExecAttach(ctx, execID.ID, container.ExecStartOptions{
		Tty: true,
	})
	if err != nil {
		return ExecInOut{}, fmt.Errorf("failed to attach exec: %v", err)
	}

	return ExecInOut{
		ExecID: execID.ID,
		Reader: resp.Reader,
		Writer: resp.Conn,
	}, nil
}

func resizeExecTTY(ctx context.Context, cli *client.Client, execID string, size [2]uint) error {
	return cli.ContainerExecResize(ctx, execID, container.ResizeOptions{
		Height: size[0],
		Width:  size[1],
	})
}
