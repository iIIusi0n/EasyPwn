package auth

import (
	"log"
	"os"
)

var globalSecretKey string

func init() {
	globalSecretKey = os.Getenv("AUTH_JWT_SECRET")
	if globalSecretKey == "" {
		log.Fatal("AUTH_JWT_SECRET is not set")
	}
}
