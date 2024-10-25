package postgres

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
	"time"
)

type MessageRepo struct {
	conn *ConnectionPool
}

func NewMessageRepo(conn *ConnectionPool) *MessageRepo {
	return &MessageRepo{conn: conn}
}

func (m *MessageRepo) CreateMessage(telegramID int64, text string, topicID int64, sender string) error {
	query := `INSERT INTO "message" (telegram_id, text, topic_id, sender) VALUES ($1, $2, $3, $4)`
	_, err := m.conn.Exec(context.Background(), query, telegramID, text, topicID, sender)
	return err
}

func (m *MessageRepo) GetMessagesByTopicID(topicID int64, after time.Time, limit int) ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT id, telegram_id, text, topic_id, sender, created_at 
				FROM "message"
				WHERE topic_id = $1 and created_at > $2
			  	ORDER BY created_at DESC
			  	LIMIT $3`
	rows, err := m.conn.Query(context.Background(), query, topicID, after, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var message models.Message
		err = rows.Scan(&message.ID, &message.TelegramID, &message.Text, &message.TopicID, &message.Sender, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
