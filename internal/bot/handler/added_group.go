package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
)

type GroupAddition struct {
	groupRepo *postgres.GroupRepo
	userRepo  *postgres.UserRepo
}

func NewGroupAddition(userRepo *postgres.UserRepo, groupRepo *postgres.GroupRepo) *GroupAddition {
	return &GroupAddition{
		groupRepo: groupRepo,
		userRepo:  userRepo,
	}
}

func (g *GroupAddition) Command() string {
	return tb.OnAddedToGroup
}

func (g *GroupAddition) Handle(ctx tb.Context) error {
	groupTelegramID := ctx.Chat().ID
	userTelegramID := ctx.Sender().ID

	user, err := g.userRepo.GetUserByTelegramID(userTelegramID)
	if err != nil {
		if postgres.IsNoRows(err) {
			return ctx.Send(NotRegisteredErrorMessage)
		}

		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	err = g.groupRepo.CreateGroup(user.ID, groupTelegramID, DefaultOpenAIModel)
	if err != nil {
		if postgres.IsUniqueViolation(err) {
			return ctx.Send(GroupAlreadyRegisteredMessage)
		}

		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	return ctx.Send(GroupRegisteredSuccessMessage)
}

func (g *GroupAddition) Middleware() []tb.MiddlewareFunc {
	return nil
}

func (g *GroupAddition) Description() string {
	return ""
}
