package service

import (
	"context"
	pb "easypwn/internal/api"
)

type ChatbotService struct {
	openAiApiKey   string
	instanceClient pb.InstanceClient

	pb.UnimplementedChatbotServer
}

func NewChatbotService(ctx context.Context, openAiApiKey string, instanceClient pb.InstanceClient) *ChatbotService {
	return &ChatbotService{
		openAiApiKey:   openAiApiKey,
		instanceClient: instanceClient,
	}
}
