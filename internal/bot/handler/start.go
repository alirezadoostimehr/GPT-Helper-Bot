package handler

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
)

type Start struct {
}

func NewStart() *Start {
	return &Start{}
}

func (s *Start) Command() string {
	return "/start"
}

func (s *Start) Handle(ctx tb.Context) error {
	return ctx.Send(fmt.Sprintf("Hey there! Is it working? Is my voice clear?"))
}

func (s *Start) Middleware() []tb.MiddlewareFunc {
	return nil
}

func (s *Start) Description() string {
	return ""
}
