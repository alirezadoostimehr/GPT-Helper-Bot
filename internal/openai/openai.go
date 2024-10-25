package openai

import (
	"context"
	openailib "github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

var GptModels = map[string]string{
	"gpt-3.5-turbo":     openailib.GPT3Dot5Turbo0125,
	openailib.GPT4:      openailib.GPT4,
	openailib.GPT4Turbo: openailib.GPT4Turbo,
	openailib.GPT4oMini: openailib.GPT4oMini,
	openailib.GPT4o:     openailib.GPT4o,
}

type Client struct {
	*openailib.Client
}

func NewGPT(apiKey string) *Client {
	return &Client{
		Client: openailib.NewClient(apiKey),
	}
}

func (g *Client) Complete(messages []map[string]string, model string) (string, error) {

	openaiMessages := make([]openailib.ChatCompletionMessage, len(messages))
	for i, message := range messages {
		role := openailib.ChatMessageRoleUser
		if message["role"] == "assistant" {
			role = openailib.ChatMessageRoleAssistant
		}
		openaiMessages[i] = openailib.ChatCompletionMessage{
			Role:    role,
			Content: message["content"],
		}
	}

	resp, err := g.CreateChatCompletion(
		context.Background(),
		openailib.ChatCompletionRequest{
			Model:    GptModels[model],
			Messages: openaiMessages,
		},
	)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (g *Client) GenerateName(prompt string) (string, error) {
	resp, err := g.CreateChatCompletion(
		context.Background(),
		openailib.ChatCompletionRequest{
			MaxTokens: 5,
			Model:     openailib.GPT4oMini,
			Messages: []openailib.ChatCompletionMessage{
				{
					Role:    openailib.ChatMessageRoleSystem,
					Content: "Generate a name with at most 6 words about the prompt",
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
