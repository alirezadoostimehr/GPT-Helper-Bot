package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	ID         int64            `json:"id"`
	TelegramID int64            `json:"telegram_id"`
	Text       string           `json:"text"`
	TopicID    int64            `json:"topic_id"`
	Sender     string           `json:"sender"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}
