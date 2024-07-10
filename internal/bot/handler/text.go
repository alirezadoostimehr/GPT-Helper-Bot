package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	tb "gopkg.in/telebot.v3"
)

type Text struct {
	openaiClient openai.Client
	mongoClient  database.MongoClient
}

func NewText(client *openai.Client, mongoClient *database.MongoClient) *Text {
	return &Text{
		openaiClient: *client,
		mongoClient:  *mongoClient,
	}
}

func (t *Text) Command() string {
	return tb.OnText
}

func (t *Text) Handle(ctx tb.Context) error {
	res, err := t.openaiClient.Complete(ctx.Text())
	if err != nil {
		return err
	}
	return ctx.Reply(res)
}

func (t *Text) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{middleware.RejectNonSupergroup(), middleware.RejectNonTopics()}
}
