package button

import tb "gopkg.in/telebot.v3"

type ButtonHandler interface {
	CallbackUnique() string
	Text() string
	Handle(ctx tb.Context) error
	Middleware() []tb.MiddlewareFunc
}
