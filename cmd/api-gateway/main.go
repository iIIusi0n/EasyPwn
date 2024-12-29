package main

import (
	"easypwn/internal/api/gateway"
	"fmt"
	"os"
)

var (
	listenPort = os.Getenv("API_LISTEN_PORT")
)

func init() {
	if listenPort == "" {
		listenPort = "8080"
	}
}

func main() {
	r := gateway.NewRouter()

	r.Run(fmt.Sprintf(":%s", listenPort))
}
