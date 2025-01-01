package data

import (
	"testing"
)

func TestGetDB(t *testing.T) {
	db := GetDB()
	if db == nil {
		t.Errorf("GetDB() returned nil")
	}

	err := db.Ping()
	if err != nil {
		t.Errorf("Failed to ping database: %v", err)
	}
}
