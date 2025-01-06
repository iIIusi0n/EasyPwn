package main

import (
	pb "easypwn/internal/api"
	"easypwn/internal/api/gateway"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	listenPort = os.Getenv("API_LISTEN_PORT")

	userListenHost = os.Getenv("USER_LISTEN_HOST")
	userListenPort = os.Getenv("USER_LISTEN_PORT")

	mailerListenHost = os.Getenv("MAILER_LISTEN_HOST")
	mailerListenPort = os.Getenv("MAILER_LISTEN_PORT")
)

func init() {
	if listenPort == "" {
		listenPort = "8080"
	}

	if userListenHost == "" || userListenPort == "" {
		log.Fatalf("USER_LISTEN_HOST and USER_LISTEN_PORT must be set")
	}

	if mailerListenHost == "" || mailerListenPort == "" {
		log.Fatalf("MAILER_LISTEN_HOST and MAILER_LISTEN_PORT must be set")
	}
}

func main() {
	userClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", userListenHost, userListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user: %v", err)
	}
	defer userClientConn.Close()

	userClient := pb.NewUserClient(userClientConn)

	mailerClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", mailerListenHost, mailerListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to mailer: %v", err)
	}
	defer mailerClientConn.Close()

	mailerClient := pb.NewMailerClient(mailerClientConn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r := gateway.NewRouter(gateway.RouterClients{
		UserClient: userClient,
		Mailer:     mailerClient,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := r.Run(fmt.Sprintf(":%s", listenPort)); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
