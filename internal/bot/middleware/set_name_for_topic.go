package middleware

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/chat"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	tb "gopkg.in/telebot.v3"
)

func SetNameForUnnamedTopic(openaiClient openai.Client) tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {

			err := next(c)
			if err != nil {
				return err
			}

			currentTopic := c.Message().ReplyTo.TopicCreated

			if currentTopic.Name == chat.DefaultTopicName {

				newName, err := openaiClient.GenerateName(c.Message().Text)
				if err != nil {
					return err
				}

				err = c.Bot().EditTopic(c.Chat(), &tb.Topic{
					Name:            newName,
					IconColor:       currentTopic.IconColor,
					IconCustomEmoji: currentTopic.IconCustomEmoji,
					ThreadID:        c.Message().ThreadID,
				})
				if err != nil {
					return err
				}
			}

			return nil
		}
	}
}
