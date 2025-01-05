package images

import "testing"

func TestEmbed(t *testing.T) {
	_, err := Dockerfiles.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
}
