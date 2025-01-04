package service

import (
	"context"
	pb "easypwn/internal/api"
	"easypwn/internal/pkg/prompt"

	openai "github.com/sashabaranov/go-openai"
)

type ChatbotService struct {
	openaiClient   *openai.Client
	instanceClient pb.InstanceClient

	pb.UnimplementedChatbotServer
}

func NewChatbotService(ctx context.Context, openAiApiKey string, instanceClient pb.InstanceClient) *ChatbotService {
	return &ChatbotService{
		openaiClient:   openai.NewClient(openAiApiKey),
		instanceClient: instanceClient,
	}
}
func (s *ChatbotService) getLogs(ctx context.Context, instanceId string) (string, error) {
	logs, err := s.instanceClient.GetInstanceLogs(ctx, &pb.GetInstanceLogsRequest{
		InstanceId: instanceId,
		Limit:      30,
	})
	if err != nil {
		return "", err
	}
	return logs.Logs, nil
}

func (s *ChatbotService) ExecuteCompletion(ctx context.Context, logs string, message string) (string, error) {
	resp, err := s.openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt.SystemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: logs,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt.LogGuidePrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func (s *ChatbotService) GetResponse(ctx context.Context, req *pb.GetResponseRequest) (*pb.GetResponseResponse, error) {
	logs, err := s.getLogs(ctx, req.InstanceId)
	if err != nil {
		return nil, err
	}
	response, err := s.ExecuteCompletion(ctx, logs, req.Message)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponseResponse{Response: response}, nil
}
