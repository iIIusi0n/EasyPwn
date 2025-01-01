package project

import (
	"testing"
)

func TestPluginTranslation(t *testing.T) {
	id, err := GetPluginIDFromName("gef")
	if err != nil {
		t.Fatal("Failed to get plugin ID: ", err)
	}
	name, err := GetPluginNameFromID(id)
	if err != nil {
		t.Fatal("Failed to get plugin name: ", err)
	}
	t.Logf("Plugin name: %s", name)
}

func TestOsTranslation(t *testing.T) {
	id, err := GetOsIDFromName("ubuntu-2410")
	if err != nil {
		t.Fatal("Failed to get os ID: ", err)
	}
	name, err := GetOsNameFromID(id)
	if err != nil {
		t.Fatal("Failed to get os name: ", err)
	}
	t.Logf("OS name: %s", name)
}
