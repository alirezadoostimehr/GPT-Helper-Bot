package handler

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/chat"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
	tb "gopkg.in/telebot.v3"
)

type NewChat struct {
	mongoClient *database.MongoClient
}

func NewNewChat(mongoClient *database.MongoClient) *NewChat {
	return &NewChat{
		mongoClient: mongoClient,
	}
}

func (n *NewChat) Command() string {
	return "/newchat"
}

func (n *NewChat) Handle(ctx tb.Context) error {
	topic, err := ctx.Bot().CreateTopic(ctx.Chat(), &tb.Topic{
		Name:            chat.DefaultTopicName,
		IconColor:       chat.DefaultTopicIconColor,
		IconCustomEmoji: chat.DefaultTopicIconCustomEmoji,
	})

	if err != nil {
		return err
	}

	superChat, err := n.mongoClient.GetSuperChatByID(context.Background(), ctx.Chat().ID)
	if err != nil {
		return err
	}

	err = n.mongoClient.CreateConversation(context.Background(), &models.Conversation{
		ChatID:      ctx.Chat().ID,
		ThreadID:    topic.ThreadID,
		OpenAIModel: superChat.OpenAIModel,
	})
	return err
}

func (n *NewChat) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{middleware.RejectNonSupergroup(), middleware.RejectNonGeneral()}
}

func (n *NewChat) Description() string {
	return "Create a new topic for chat"
}
