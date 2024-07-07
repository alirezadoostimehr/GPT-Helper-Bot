package handler

import tb "gopkg.in/telebot.v3"

type Command interface {
	Command() string
	Handle(ctx tb.Context) error
}
