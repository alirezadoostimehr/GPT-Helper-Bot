package gpt

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type GPT struct {
	*openai.Client
}

func NewGPT(apiKey string) *GPT {
	return &GPT{
		Client: openai.NewClient(apiKey),
	}
}

func (g *GPT) Complete(prompt string) (string, error) {
	resp, err := g.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
