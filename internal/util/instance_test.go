package util

import (
	"testing"
)

func TestDockerfileToImageName(t *testing.T) {
	tests := []struct {
		dockerfile string
		want       string
	}{
		{
			dockerfile: "Dockerfile.ubuntu-2410.gef",
			want:       "easypwn/ubuntu-2410/gef",
		},
		{
			dockerfile: "Dockerfile.ubuntu-2410.pwndbg",
			want:       "easypwn/ubuntu-2410/pwndbg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.dockerfile, func(t *testing.T) {
			if got := DockerfileToImageName(tt.dockerfile); got != tt.want {
				t.Errorf("DockerfileToImageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
