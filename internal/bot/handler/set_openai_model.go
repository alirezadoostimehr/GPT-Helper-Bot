package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	tb "gopkg.in/telebot.v3"
)

type SetOpenAIModel struct {
}

func NewSetOpenAIModel() *SetOpenAIModel {
	return &SetOpenAIModel{}
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
		middleware.RejectNonSupergroup(),
		middleware.RejectNonGeneral(),
	}
}

func (h *SetOpenAIModel) Description() string {
	return "Set OpenAI model for the topic"
}

func CreateButtons() [][]tb.InlineButton {
	var buttons [][]tb.InlineButton
	for model := range openai.GptModels {
		buttons = append(buttons, []tb.InlineButton{
			{
				Unique: model,
				Text:   model,
			},
		})
	}
	return buttons
}
