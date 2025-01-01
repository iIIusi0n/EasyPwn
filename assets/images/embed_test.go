package images

import "testing"

func TestEmbed(t *testing.T) {
	files, err := Dockerfiles.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Logf("Found file: %s", file.Name())
	}
}
