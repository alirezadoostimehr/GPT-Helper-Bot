package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
)

type Start struct {
	userRepo *postgres.UserRepo
}

func NewStart(repo *postgres.UserRepo) *Start {
	return &Start{
		userRepo: repo,
	}
}

func (s *Start) Command() string {
	return "/start"
}

func (s *Start) Handle(ctx tb.Context) error {
	telegramID := ctx.Sender().ID

	err := s.userRepo.CreateUser(telegramID)
	if err != nil {
		if postgres.IsUniqueViolation(err) {
			return ctx.Send(UserAlreadyRegisteredMessage)
		}

		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	return ctx.Send(UserRegisteredSuccessMessage)
}

func (s *Start) Middleware() []tb.MiddlewareFunc {
	return nil
}

func (s *Start) Description() string {
	return ""
}
