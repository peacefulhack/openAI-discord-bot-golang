package utils

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var AIToken string

func NewAI() error {
	res, err := GetTokenEnv("OPENAI_TOKEN")
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	AIToken = res
	return nil
}

func GPT3TurboChat(message string) (string, error) {
	client := openai.NewClient(AIToken)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	})

	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}

	return resp.Choices[0].Message.Content, nil
}

func DallEChat(message string) (string, error) {
	client := openai.NewClient(AIToken)

	res, err := client.CreateImage(context.Background(), openai.ImageRequest{
		Prompt: message,
		N:      1,
		Size:   "256x256",
	})

	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}

	return res.Data[0].URL, nil
}
