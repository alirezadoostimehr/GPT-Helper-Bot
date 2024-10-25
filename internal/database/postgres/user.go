package postgres

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
)

type UserRepo struct {
	conn *ConnectionPool
}

func NewUserRepo(conn *ConnectionPool) *UserRepo {
	return &UserRepo{conn: conn}
}

func (u *UserRepo) CreateUser(telegramID int64) error {
	query := `
		INSERT INTO "user" (telegram_id)
		VALUES ($1)
	`
	_, err := u.conn.Exec(context.Background(), query, telegramID)
	return err
}

func (u *UserRepo) GetUserByTelegramID(telegramID int64) (models.User, error) {
	query := `
		SELECT id, telegram_id, created_at
		FROM "user"
		WHERE telegram_id = $1
	`
	var user models.User
	err := u.conn.QueryRow(context.Background(), query, telegramID).Scan(&user.ID, &user.TelegramID, &user.CreatedAt)
	return user, err
}
