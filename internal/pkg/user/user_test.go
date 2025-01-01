package user

import (
	"context"
	"easypwn/internal/data"
	"testing"
)

func TestUser(t *testing.T) {
	u, err := NewUser(context.Background(), data.GetDB(), "test-user", "test-password", "test-email")
	if err != nil {
		t.Fatal("Failed to create user: ", err)
	}
	defer u.Delete(context.Background(), data.GetDB())

	t.Logf("User created: %+v", u)
}
