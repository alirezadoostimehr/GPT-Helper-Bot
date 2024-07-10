package database

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	*mongo.Client
	Database *mongo.Database
}

func NewMongoClient() (*MongoClient, error) {
	mc, err := newClient()
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client:   mc,
		Database: mc.Database(config.GlobalConfig.Mongo.NAME),
	}, nil
}

func newClient() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.GlobalConfig.Mongo.URI))
	if err != nil {
		return nil, err
	}

	return client, nil
}

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
