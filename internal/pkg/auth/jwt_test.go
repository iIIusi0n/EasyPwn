package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestEncode(t *testing.T) {
	token := NewToken(uuid.New().String(), "test@test.com")
	_, err := token.Encode()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecode(t *testing.T) {
	token := NewToken(uuid.New().String(), "test@test.com")
	encoded, err := token.Encode()
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.UserID != token.UserID || decoded.Email != token.Email {
		t.Fatal("decoded token does not match original token")
	}
	if decoded.IsExpired() {
		t.Fatal("decoded token is expired")
	}
}
