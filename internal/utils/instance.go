package utils

import (
	"archive/tar"
	"bytes"
	"fmt"
	"strings"
)

func DockerfileToImageName(dockerfile string) string {
	words := strings.Split(dockerfile, ".")
	return fmt.Sprintf("easypwn/%s/%s", words[1], words[2])
}

func CreateDockerfileTar(filename string, content []byte) (*bytes.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	tarHeader := &tar.Header{
		Name: filename,
		Size: int64(len(content)),
	}

	if err := tw.WriteHeader(tarHeader); err != nil {
		return nil, err
	}

	if _, err := tw.Write(content); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}
