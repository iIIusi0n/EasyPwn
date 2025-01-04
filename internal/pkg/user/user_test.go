package user

import (
	"context"
	"easypwn/internal/data"
	"testing"
)

func TestUser(t *testing.T) {
	u, err := NewUser(context.Background(), data.GetDB(), "test-email", "test-password")
	if err != nil {
		t.Fatal("Failed to create user: ", err)
	}
	defer u.Delete(context.Background(), data.GetDB())

	t.Logf("User created: %+v", u)
}
