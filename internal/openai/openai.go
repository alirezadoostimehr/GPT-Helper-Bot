package openai

import (
	"context"
	openailib "github.com/sashabaranov/go-openai"
)

type Client struct {
	*openailib.Client
}

func NewGPT(apiKey string) *Client {
	return &Client{
		Client: openailib.NewClient(apiKey),
	}
}

func (g *Client) Complete(prompt string) (string, error) {
	resp, err := g.CreateChatCompletion(
		context.Background(),
		openailib.ChatCompletionRequest{
			Model: openailib.GPT4Turbo,
			Messages: []openailib.ChatCompletionMessage{
				{
					Role:    openailib.ChatMessageRoleUser,
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
