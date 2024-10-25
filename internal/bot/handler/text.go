package handler

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/middleware"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"
	"sort"
	"time"
)

type Text struct {
	openaiClient openai.Client
	topicRepo    *postgres.TopicRepo
	messageRepo  *postgres.MessageRepo
}

func NewText(client *openai.Client, topicRepo *postgres.TopicRepo, messageRepo *postgres.MessageRepo) *Text {
	return &Text{
		openaiClient: *client,
		topicRepo:    topicRepo,
		messageRepo:  messageRepo,
	}
}

func (t *Text) Command() string {
	return tb.OnText
}

func (t *Text) Handle(ctx tb.Context) error {
	threadID := ctx.Message().ThreadID

	topic, err := t.topicRepo.GetTopicByThreadID(threadID)
	if err != nil {
		log.Error(err)
		return ctx.Reply(InternalErrorMessage)
	}

	messages, err := t.messageRepo.GetMessagesByTopicID(topic.ID, time.Now().Add(-24*time.Hour), 30)
	if err != nil {
		log.Error(err)
		return ctx.Reply(InternalErrorMessage)
	}

	conversationMessages := CreateConversationFromMessages(messages, []string{ctx.Text()})
	res, err := t.openaiClient.Complete(conversationMessages, topic.OpenAIModel)
	if err != nil {
		log.Error(err)
		return ctx.Reply(InternalErrorMessage)
	}

	err = t.messageRepo.CreateMessage(int64(ctx.Message().ID), ctx.Message().Text, topic.ID, "user")
	if err != nil {
		log.Error(err)
		return ctx.Reply(InternalErrorMessage)
	}

	sentMessage, err := ctx.Bot().Reply(ctx.Message(), res)
	if err != nil {
		log.Error(err)
		return ctx.Reply(InternalErrorMessage)
	}
	err = t.messageRepo.CreateMessage(int64(sentMessage.ID), res, topic.ID, "assistant")
	return err
}

func (t *Text) Middleware() []tb.MiddlewareFunc {
	return []tb.MiddlewareFunc{
		middleware.RejectNonSupergroup(),
		middleware.RejectNonTopics(),
		middleware.SetNameForUnnamedTopic(t.openaiClient),
	}
}

func (t *Text) Description() string {
	return ""
}

func CreateConversationFromMessages(messages []models.Message, additionalMessage []string) []map[string]string {
	sort.SliceStable(messages, func(i, j int) bool {
		return messages[i].CreatedAt.Time.Before(messages[j].CreatedAt.Time)
	})

	res := make([]map[string]string, 0)
	for _, message := range messages {
		if message.Sender == "user" {
			res = append(res, map[string]string{"role": "user", "content": message.Text})
		} else {
			res = append(res, map[string]string{"role": "assistant", "content": message.Text})
		}
	}
	for _, message := range additionalMessage {
		res = append(res, map[string]string{"role": "user", "content": message})
	}
	return res
}
