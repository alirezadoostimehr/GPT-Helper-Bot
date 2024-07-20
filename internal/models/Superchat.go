package models

const (
	SuperChatCollectionName = "super_chats"
)

type SuperChat struct {
	ChatID      int64  `bson:"chat_id"`
	OpenAIModel string `bson:"model" default:"gp-3.5-turbo"`
}
