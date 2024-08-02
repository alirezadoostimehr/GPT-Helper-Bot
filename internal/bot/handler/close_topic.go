package handler

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database"
	tb "gopkg.in/telebot.v3"
)

type CloseTopic struct {
	mongoClient *database.MongoClient
}

func NewCloseTopic(mongoClient *database.MongoClient) *CloseTopic {
	return &CloseTopic{
		mongoClient: mongoClient,
	}
}

func (n *CloseTopic) Command() string {
	return "/close_topic"
}

func (n *CloseTopic) Handle(ctx tb.Context) error {
	err := ctx.Bot().DeleteTopic(ctx.Chat(), &tb.Topic{ThreadID: ctx.Message().ThreadID})
	if err != nil {
		return err
	}

	err = n.mongoClient.DeleteConversation(context.Background(), ctx.Chat().ID, ctx.Message().ThreadID)

	return err
}

func (n *CloseTopic) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{
		middleware.RejectNonSupergroup(),
		middleware.RejectNonTopics(),
	}
}

func (n *CloseTopic) Description() string {
	return "Close the current topic"
}
