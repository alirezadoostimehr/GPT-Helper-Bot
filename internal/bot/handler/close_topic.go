package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
)

type CloseTopic struct {
	topicRepo *postgres.TopicRepo
}

func NewCloseTopic(topicRepo *postgres.TopicRepo) *CloseTopic {
	return &CloseTopic{
		topicRepo: topicRepo,
	}
}

func (c *CloseTopic) Command() string {
	return "/close_topic"
}

func (c *CloseTopic) Handle(ctx tb.Context) error {
	err := ctx.Bot().DeleteTopic(ctx.Chat(), &tb.Topic{ThreadID: ctx.Message().ThreadID})
	if err != nil {
		log.Error(err)
		return err
	}

	threadID := ctx.Message().ThreadID
	err = c.topicRepo.DeleteTopicByThreadID(threadID)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (c *CloseTopic) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{
		middleware.RejectNonSupergroup(),
		middleware.RejectNonTopics(),
	}
}

func (c *CloseTopic) Description() string {
	return "Close the current topic"
}
