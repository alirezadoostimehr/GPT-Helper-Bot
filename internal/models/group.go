package models

import "github.com/jackc/pgx/v5/pgtype"

type Group struct {
	ID          int64            `json:"id"`
	UserID      int64            `json:"user_id"`
	TelegramID  int64            `json:"telegram_id"`
	OpenAIModel string           `json:"openai_model"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}
