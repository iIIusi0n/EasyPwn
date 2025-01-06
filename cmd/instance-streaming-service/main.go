package main

import (
	"fmt"
	"log"
	"os"

	pb "easypwn/internal/api"
	"easypwn/internal/api/stream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	listenPort = os.Getenv("INSTANCE_STREAMING_SERVICE_LISTEN_PORT")

	instanceListenHost = os.Getenv("INSTANCE_LISTEN_HOST")
	instanceListenPort = os.Getenv("INSTANCE_LISTEN_PORT")

	projectListenHost = os.Getenv("PROJECT_LISTEN_HOST")
	projectListenPort = os.Getenv("PROJECT_LISTEN_PORT")
)

func init() {
	if listenPort == "" {
		listenPort = "50055"
	}

	if instanceListenHost == "" || instanceListenPort == "" {
		log.Fatalf("instance listen host or port is not set")
	}

	if projectListenHost == "" || projectListenPort == "" {
		log.Fatalf("project listen host or port is not set")
	}
}

func main() {
	instanceClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", instanceListenHost, instanceListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to instance: %v", err)
	}
	defer instanceClientConn.Close()

	instanceClient := pb.NewInstanceClient(instanceClientConn)

	projectClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", projectListenHost, projectListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to project: %v", err)
	}
	defer projectClientConn.Close()

	projectClient := pb.NewProjectClient(projectClientConn)

	r := stream.NewRouter(stream.RouterClients{
		ProjectClient:  projectClient,
		InstanceClient: instanceClient,
	})

	if err := r.Run(fmt.Sprintf(":%s", listenPort)); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
