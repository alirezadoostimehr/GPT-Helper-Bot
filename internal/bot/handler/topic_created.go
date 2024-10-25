package handler

import (
	"fmt"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
)

type TopicCreated struct {
	topicRepo *postgres.TopicRepo
}

func NewTopicCreated(topicRepo *postgres.TopicRepo) *TopicCreated {
	return &TopicCreated{
		topicRepo: topicRepo,
	}
}

func (t *TopicCreated) Command() string {
	return tb.OnTopicCreated
}

func (t *TopicCreated) Handle(ctx tb.Context) error {
	topic, err := t.topicRepo.GetTopicByThreadID(ctx.Message().ThreadID)
	if err != nil {
		log.Error(err)
		return ctx.Send(InternalErrorMessage)
	}

	return ctx.Reply(fmt.Sprintf(TopicCreationSuccessMessage, topic.OpenAIModel))
}

func (t *TopicCreated) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{}
}
