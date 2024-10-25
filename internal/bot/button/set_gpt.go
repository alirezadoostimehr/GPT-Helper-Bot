package button

import (
	"fmt"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
)

type SetModel struct {
	groupRepo *postgres.GroupRepo
	modelName string
}

func NewSetModel(groupRepo *postgres.GroupRepo, model string) *SetModel {
	return &SetModel{
		groupRepo: groupRepo,
		modelName: model,
	}
}

func (s SetModel) CallbackUnique() string {
	return "\f" + s.modelName
}

func (s SetModel) Text() string {
	return s.modelName
}

func (s SetModel) Handle(ctx tb.Context) error {
	chatID := ctx.Chat().ID
	group, err := s.groupRepo.GetGroupByTelegramID(chatID)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.groupRepo.SetGroupOpenAIModel(group.TelegramID, s.modelName)
	if err != nil {
		log.Error(err)
		return err
	}

	return ctx.Reply(fmt.Sprintf("%s model is set", s.modelName))
}

func (s SetModel) Middleware() []tb.MiddlewareFunc {
	return nil
}
