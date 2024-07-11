package models

const (
	ConversationCollectionName = "conversations"
)

type Conversation struct {
	Messages    []string `bson:"messages"`
	ChatID      int64    `bson:"chat_id"`
	ThreadID    int      `bson:"thread_id"`
	OpenAIModel string   `bson:"openai_model"`
}
