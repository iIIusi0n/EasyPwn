package project

import (
	"testing"
)

func TestPluginTranslation(t *testing.T) {
	id, err := GetPluginIDFromName("gef")
	if err != nil {
		t.Fatal("Failed to get plugin ID: ", err)
	}
	_, err = GetPluginNameFromID(id)
	if err != nil {
		t.Fatal("Failed to get plugin name: ", err)
	}
}

func TestOsTranslation(t *testing.T) {
	id, err := GetOsIDFromName("ubuntu-2410")
	if err != nil {
		t.Fatal("Failed to get os ID: ", err)
	}
	_, err = GetOsNameFromID(id)
	if err != nil {
		t.Fatal("Failed to get os name: ", err)
	}
}
