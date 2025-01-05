package main

import (
	pb "easypwn/internal/api"
	"easypwn/internal/api/gateway"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	listenPort = os.Getenv("API_LISTEN_PORT")

	userListenHost = os.Getenv("USER_LISTEN_HOST")
	userListenPort = os.Getenv("USER_LISTEN_PORT")
)

func init() {
	if listenPort == "" {
		listenPort = "8080"
	}

	if userListenHost == "" || userListenPort == "" {
		log.Fatalf("USER_LISTEN_HOST and USER_LISTEN_PORT must be set")
	}
}

func main() {
	userClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", userListenHost, userListenPort))
	if err != nil {
		log.Fatalf("failed to connect to user: %v", err)
	}
	defer userClientConn.Close()

	userClient := pb.NewUserClient(userClientConn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r := gateway.NewRouter(gateway.RouterClients{
		UserClient: userClient,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := r.Run(fmt.Sprintf(":%s", listenPort)); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
