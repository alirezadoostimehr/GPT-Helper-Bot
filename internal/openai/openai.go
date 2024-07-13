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

func (g *Client) Complete(messages []string) (string, error) {
	openaiMessages := make([]openailib.ChatCompletionMessage, len(messages))
	for i, message := range messages {
		role := openailib.ChatMessageRoleUser
		if i%2 == 1 {
			role = openailib.ChatMessageRoleAssistant
		}
		openaiMessages[i] = openailib.ChatCompletionMessage{
			Role:    role,
			Content: message,
		}
	}

	resp, err := g.CreateChatCompletion(
		context.Background(),
		openailib.ChatCompletionRequest{
			Model:    openailib.GPT3Dot5Turbo0125,
			Messages: openaiMessages,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (g *Client) GenerateName(prompt string) (string, error) {
	resp, err := g.CreateChatCompletion(
		context.Background(),
		openailib.ChatCompletionRequest{
			MaxTokens: 5,
			Model:     openailib.GPT3Dot5Turbo,
			Messages: []openailib.ChatCompletionMessage{
				{
					Role:    openailib.ChatMessageRoleSystem,
					Content: "Generate a name with at most 3 words about the prompt",
				},
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
