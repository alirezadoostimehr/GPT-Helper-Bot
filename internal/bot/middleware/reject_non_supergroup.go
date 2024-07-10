package middleware

import (
	tb "gopkg.in/telebot.v3"
)

func RejectNonSupergroup() tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			if c.Chat().Type != tb.ChatSuperGroup {
				return c.Reply("This command is only supported in supergroups.")
			}
			return next(c)
		}
	}
}
