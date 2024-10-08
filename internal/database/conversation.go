package database

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
)

func (mc *MongoClient) CreateConversation(ctx context.Context, conversation *models.Conversation) error {
	_, err := mc.Database.Collection(models.ConversationCollectionName).InsertOne(ctx, conversation)
	return err
}

func (mc *MongoClient) GetConversationByIDs(ctx context.Context, chatID int64, threadID int) (*models.Conversation, error) {
	var conversation models.Conversation
	err := mc.Database.Collection(models.ConversationCollectionName).FindOne(ctx, map[string]any{"chat_id": chatID, "thread_id": threadID}).Decode(&conversation)
	if err != nil {
		return nil, err
	}

	return &conversation, nil
}

func (mc *MongoClient) UpdateConversation(ctx context.Context, conversation *models.Conversation) error {
	_, err := mc.Database.Collection(models.ConversationCollectionName).UpdateOne(ctx, map[string]any{"chat_id": conversation.ChatID, "thread_id": conversation.ThreadID}, map[string]interface{}{"$set": conversation})
	return err
}

func (mc *MongoClient) DeleteConversation(ctx context.Context, chatID int64, threadID int) error {
	_, err := mc.Database.Collection(models.ConversationCollectionName).DeleteOne(ctx, map[string]any{"chat_id": chatID, "thread_id": threadID})
	return err
}
