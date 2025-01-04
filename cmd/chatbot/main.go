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
	openAiApiKey = os.Getenv("CHATBOT_OPENAI_API_KEY")
	listenPort   = os.Getenv("CHATBOT_LISTEN_PORT")

	instanceListenHost = os.Getenv("INSTANCE_LISTEN_HOST")
	instanceListenPort = os.Getenv("INSTANCE_LISTEN_PORT")
)

func init() {
	if listenPort == "" {
		listenPort = "50054"
	}

	if instanceListenHost == "" || instanceListenPort == "" {
		log.Fatalf("instance listen host or port is not set")
	}

	if openAiApiKey == "" {
		log.Fatalf("openai api key is not set")
	}
}

func main() {
	instanceClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", instanceListenHost, instanceListenPort))
	if err != nil {
		log.Fatalf("failed to connect to instance: %v", err)
	}
	defer instanceClientConn.Close()

	instanceClient := pb.NewInstanceClient(instanceClientConn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatbotServer(s, service.NewChatbotService(context.Background(), openAiApiKey, instanceClient))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
