package middleware

import tb "gopkg.in/telebot.v3"

func RejectPrivateChat() tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			if c.Chat().Type == tb.ChatPrivate {
				return c.Reply("Only message in group chat is supported.")
			}
			return next(c)
		}
	}
}
