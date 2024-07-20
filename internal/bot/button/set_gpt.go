package button

import (
	"context"
	"fmt"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database"
	tb "gopkg.in/telebot.v3"
)

type SetModel struct {
	mongoClient *database.MongoClient
	modelName   string
}

func NewSetModel(mongoClient *database.MongoClient, model string) *SetModel {
	return &SetModel{
		mongoClient: mongoClient,
		modelName:   model,
	}
}

func (s SetModel) CallbackUnique() string {
	return "\f" + s.modelName
}

func (s SetModel) Text() string {
	return s.modelName
}

func (s SetModel) Handle(ctx tb.Context) error {
	superChat, err := s.mongoClient.GetSuperChatOrCreate(context.Background(), ctx.Chat().ID)
	if err != nil {
		return err
	}

	superChat.OpenAIModel = s.modelName
	err = s.mongoClient.UpdateSuperChat(context.Background(), superChat)
	if err != nil {
		return err
	}

	return ctx.Reply(fmt.Sprintf("%s model is set", s.modelName))
}

func (s SetModel) Middleware() []tb.MiddlewareFunc {
	return nil
}
