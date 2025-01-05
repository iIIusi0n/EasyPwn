package instance

import (
	"context"
	"testing"
)

func TestInitImages(t *testing.T) {
	imageNames, err := getImageNames(context.Background(), cli)
	if err != nil {
		t.Fatal("Failed to get image names: ", err)
	}

	for _, imageName := range imageNames {
		if imageName == "easypwn/ubuntu-2410:gef" {
			return
		}
	}

	t.Fatal("Image not found")
}
