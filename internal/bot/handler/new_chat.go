package handler

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
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
		Name:            "New GPT chat",
		IconColor:       1,
		IconCustomEmoji: "😮‍💨",
	})

	if err != nil {
		return err
	}

	err = n.mongoClient.CreateConversation(context.Background(), &models.Conversation{
		ChatID:   ctx.Chat().ID,
		ThreadID: topic.ThreadID,
	})
	return err
}

func (n *NewChat) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{middleware.RejectNonSupergroup(), middleware.RejectNonGeneral()}
}
