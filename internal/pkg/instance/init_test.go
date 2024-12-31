package instance

import (
	"easypwn/assets/images"
	"testing"
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
