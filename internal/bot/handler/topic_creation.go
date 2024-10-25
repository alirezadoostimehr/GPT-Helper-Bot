package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/chat"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
)

type TopicCreation struct {
	groupRepo *postgres.GroupRepo
	topicRepo *postgres.TopicRepo
}

func NewTopicCreation(groupRepo *postgres.GroupRepo, topicRepo *postgres.TopicRepo) *TopicCreation {
	return &TopicCreation{
		groupRepo: groupRepo,
		topicRepo: topicRepo,
	}
}

func (t *TopicCreation) Command() string {
	return "/newtopic"
}

func (t *TopicCreation) Handle(ctx tb.Context) error {
	topic, err := ctx.Bot().CreateTopic(ctx.Chat(), &tb.Topic{
		Name: chat.DefaultTopicName,
	})

	if err != nil {
		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	groupTelegramID := ctx.Chat().ID
	group, err := t.groupRepo.GetGroupByTelegramID(groupTelegramID)
	if err != nil {
		if postgres.IsNoRows(err) {
			return ctx.Send(GroupNotRegisteredErrorMessage)
		}

		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	err = t.topicRepo.CreateTopic(int64(topic.ThreadID), group.ID, chat.DefaultTopicName, group.OpenAIModel)
	if err != nil {
		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	return ctx.Bot().React(ctx.Chat(), ctx.Message(), tb.ReactionOptions{Reactions: []tb.Reaction{ReactionSuccess}})
}

func (t *TopicCreation) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{middleware.RejectNonSupergroup(), middleware.RejectNonGeneral()}
}
