package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "easypwn/internal/api"
	"easypwn/internal/api/stream"

	"google.golang.org/grpc"
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
	instanceClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", instanceListenHost, instanceListenPort))
	if err != nil {
		log.Fatalf("failed to connect to instance: %v", err)
	}
	defer instanceClientConn.Close()

	instanceClient := pb.NewInstanceClient(instanceClientConn)

	projectClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", projectListenHost, projectListenPort))
	if err != nil {
		log.Fatalf("failed to connect to project: %v", err)
	}
	defer projectClientConn.Close()

	projectClient := pb.NewProjectClient(projectClientConn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r := stream.NewRouter(stream.RouterClients{
		ProjectClient:  projectClient,
		InstanceClient: instanceClient,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := r.Run(fmt.Sprintf(":%s", listenPort)); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
