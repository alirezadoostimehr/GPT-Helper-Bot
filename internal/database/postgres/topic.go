package postgres

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
)

type TopicRepo struct {
	conn *ConnectionPool
}

func NewTopicRepo(conn *ConnectionPool) *TopicRepo {
	return &TopicRepo{conn: conn}
}

func (t *TopicRepo) CreateTopic(threadID int64, groupID int64, name string, openaiModel string) error {
	query := `INSERT INTO "topic" (thread_id, group_id, name, openai_model) VALUES ($1, $2, $3, $4)`
	_, err := t.conn.Exec(context.Background(), query, threadID, groupID, name, openaiModel)
	return err
}

func (t *TopicRepo) GetTopicByThreadID(threadID int) (models.Topic, error) {
	var topic models.Topic
	query := `SELECT id, thread_id, group_id, name, openai_model, created_at FROM "topic" WHERE thread_id = $1`
	err := t.conn.QueryRow(context.Background(), query, threadID).Scan(&topic.ID, &topic.ThreadID, &topic.GroupID, &topic.Name, &topic.OpenAIModel, &topic.CreatedAt)
	return topic, err
}

func (t *TopicRepo) DeleteTopicByThreadID(threadID int) error {
	query := `DELETE FROM "topic" WHERE thread_id = $1`
	_, err := t.conn.Exec(context.Background(), query, threadID)
	return err
}
