package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "easypwn/internal/api"
	"easypwn/internal/service"

	"google.golang.org/grpc"
)

var (
	listenPort = os.Getenv("PROJECT_LISTEN_PORT")
)

func init() {
	if listenPort == "" {
		listenPort = "50052"
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProjectServer(s, service.NewProjectService(context.Background()))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
