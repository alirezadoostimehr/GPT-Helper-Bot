package models

import "github.com/jackc/pgx/v5/pgtype"

type Topic struct {
	ID          int64            `json:"id"`
	ThreadID    int              `json:"thread_id"`
	GroupID     int64            `json:"group_id"`
	Name        string           `json:"name"`
	OpenAIModel string           `json:"openai_model"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}
