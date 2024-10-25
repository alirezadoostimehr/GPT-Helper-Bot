package postgres

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
)

type GroupRepo struct {
	conn *ConnectionPool
}

func NewGroupRepo(conn *ConnectionPool) *GroupRepo {
	return &GroupRepo{conn: conn}
}

func (g *GroupRepo) CreateGroup(userID int64, telegramID int64, openAIModel string) error {
	query := `INSERT INTO "group" (user_id, telegram_id, openai_model) VALUES ($1, $2, $3)`
	_, err := g.conn.Exec(context.Background(), query, userID, telegramID, openAIModel)
	return err
}

func (g *GroupRepo) GetGroupByTelegramID(telegramID int64) (models.Group, error) {
	var group models.Group
	query := `SELECT id, user_id, telegram_id, openai_model, created_at FROM "group" WHERE telegram_id = $1`
	err := g.conn.QueryRow(context.Background(), query, telegramID).Scan(&group.ID, &group.UserID, &group.TelegramID, &group.OpenAIModel, &group.CreatedAt)
	return group, err
}

func (g *GroupRepo) SetGroupOpenAIModel(telegramID int64, openAIModel string) error {
	query := `UPDATE "group" SET openai_model = $1 WHERE telegram_id = $2`
	_, err := g.conn.Exec(context.Background(), query, openAIModel, telegramID)
	return err
}
