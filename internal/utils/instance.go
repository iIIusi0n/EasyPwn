package utils

import (
	"fmt"
	"strings"
)

func DockerfileToImageName(dockerfile string) string {
	words := strings.Split(dockerfile, ".")
	return fmt.Sprintf("easypwn/%s/%s", words[1], words[2])
}
