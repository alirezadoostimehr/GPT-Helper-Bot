package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/gpt"
	tb "gopkg.in/telebot.v3"
)

type Text struct {
	Gptbot gpt.GPT
}

func NewText(gptbot *gpt.GPT) *Text {
	return &Text{
		Gptbot: *gptbot,
	}
}

func (t *Text) Command() string {
	return tb.OnText
}

func (t *Text) Handle(ctx tb.Context) error {
	res, err := t.Gptbot.Complete(ctx.Text())
	if err != nil {
		return err
	}
	return ctx.Send(res)
}
