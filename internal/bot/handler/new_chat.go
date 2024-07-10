package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	tb "gopkg.in/telebot.v3"
)

type NewChat struct{}

func NewNewChat() *NewChat {
	return &NewChat{}
}

func (n *NewChat) Command() string {
	return "/newchat"
}

func (n *NewChat) Handle(ctx tb.Context) error {
	_, err := ctx.Bot().CreateTopic(ctx.Chat(), &tb.Topic{
		Name:            "New GPT chat",
		IconColor:       1,
		IconCustomEmoji: "😮‍💨",
	})
	return err
}

func (n *NewChat) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{middleware.RejectNonSupergroup(), middleware.RejectNonGeneral()}
}
