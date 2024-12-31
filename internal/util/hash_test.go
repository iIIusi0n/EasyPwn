package util

import "testing"

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{
			name:     "empty password",
			password: "",
			want:     "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "simple password",
			password: "password123",
			want:     "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
		},
		{
			name:     "complex password",
			password: "P@ssw0rd!123",
			want:     "c84e1a4968c055bcadba3ca783e5887b12a623c4e8c6b538b8090fa9cf18d71f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashPassword(tt.password); got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
