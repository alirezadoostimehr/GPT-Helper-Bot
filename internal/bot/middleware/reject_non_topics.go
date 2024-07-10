package middleware

import (
	tb "gopkg.in/telebot.v3"
)

func RejectNonTopics() tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			if c.Message().ThreadID == 0 {
				return c.Reply("This command is only supported in non general topics.")
			}
			return next(c)
		}
	}
}
