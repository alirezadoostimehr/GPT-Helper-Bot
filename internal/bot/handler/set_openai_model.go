package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	tb "gopkg.in/telebot.v3"
)

type SetOpenAIModel struct {
	mongoClient *database.MongoClient
}

func NewSetOpenAIModel(mongoClient *database.MongoClient) *SetOpenAIModel {
	return &SetOpenAIModel{
		mongoClient: mongoClient,
	}
}

func (h *SetOpenAIModel) Command() string {
	return "/choose_openai_model"
}

func (h *SetOpenAIModel) Handle(ctx tb.Context) error {
	buttons := CreateButtons()
	return ctx.Reply("Choose ", &tb.ReplyMarkup{
		InlineKeyboard: buttons,
	})
}

func (h *SetOpenAIModel) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{
		//middleware.RejectNonSupergroup(),
		//middleware.RejectNonGeneral(),
	}
}

func (h *SetOpenAIModel) Description() string {
	return "Set OpenAI model for the topic"
}

func CreateButtons() [][]tb.InlineButton {
	var buttons [][]tb.InlineButton
	for _, model := range openai.GptModels {
		buttons = append(buttons, []tb.InlineButton{
			tb.InlineButton{
				Unique: model,
				Text:   model,
			},
		})
	}
	return buttons
}
