package models

const (
	ConversationCollectionName = "conversations"
)

type Conversation struct {
	Dialog      []string `bson:"dialog"`
	ChatID      int64    `bson:"chat_id"`
	ThreadID    int      `bson:"thread_id"`
	OpenAIModel string   `bson:"openai_model"`
}
