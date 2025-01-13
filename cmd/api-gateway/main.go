package main

import (
	pb "easypwn/internal/api"
	"easypwn/internal/api/gateway"
	"fmt"
	"log"
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

	projectListenHost = os.Getenv("PROJECT_LISTEN_HOST")
	projectListenPort = os.Getenv("PROJECT_LISTEN_PORT")

	instanceListenHost = os.Getenv("INSTANCE_LISTEN_HOST")
	instanceListenPort = os.Getenv("INSTANCE_LISTEN_PORT")

	chatbotListenHost = os.Getenv("CHATBOT_LISTEN_HOST")
	chatbotListenPort = os.Getenv("CHATBOT_LISTEN_PORT")

	dockerHostMountPath = os.Getenv("DOCKER_HOST_MOUNT_PATH")
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

	if projectListenHost == "" || projectListenPort == "" {
		log.Fatalf("PROJECT_LISTEN_HOST and PROJECT_LISTEN_PORT must be set")
	}

	if instanceListenHost == "" || instanceListenPort == "" {
		log.Fatalf("INSTANCE_LISTEN_HOST and INSTANCE_LISTEN_PORT must be set")
	}

	if chatbotListenHost == "" || chatbotListenPort == "" {
		log.Fatalf("CHATBOT_LISTEN_HOST and CHATBOT_LISTEN_PORT must be set")
	}

	if dockerHostMountPath == "" {
		log.Fatalf("DOCKER_HOST_MOUNT_PATH must be set")
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

	projectClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", projectListenHost, projectListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to project: %v", err)
	}
	defer projectClientConn.Close()

	projectClient := pb.NewProjectClient(projectClientConn)

	instanceClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", instanceListenHost, instanceListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to instance: %v", err)
	}
	defer instanceClientConn.Close()

	instanceClient := pb.NewInstanceClient(instanceClientConn)

	chatbotClientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", chatbotListenHost, chatbotListenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to chatbot: %v", err)
	}
	defer chatbotClientConn.Close()

	chatbotClient := pb.NewChatbotClient(chatbotClientConn)

	r := gateway.NewRouter(gateway.RouterClients{
		UserClient:     userClient,
		Mailer:         mailerClient,
		ProjectClient:  projectClient,
		InstanceClient: instanceClient,
		ChatbotClient:  chatbotClient,
	})

	if err := r.Run(fmt.Sprintf(":%s", listenPort)); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
