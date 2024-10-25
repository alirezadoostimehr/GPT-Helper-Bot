package models

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID         int64            `json:"id"`
	TelegramID int64            `json:"telegram_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}
