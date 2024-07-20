package database

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/models"
	"github.com/pkg/errors"
)

func (mc *MongoClient) CreateSuperChat(ctx context.Context, superChat *models.SuperChat) error {
	_, err := mc.Database.Collection(models.SuperChatCollectionName).InsertOne(ctx, superChat)
	return err
}

func (mc *MongoClient) GetSuperChatOrCreate(ctx context.Context, chatID int64) (*models.SuperChat, error) {
	superChat, err := mc.GetSuperChatByID(ctx, chatID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if superChat == nil {
		superChat = &models.SuperChat{
			ChatID: chatID,
		}
		err = mc.CreateSuperChat(ctx, superChat)
		if err != nil {
			return nil, err
		}
	}

	return superChat, nil

}

func (mc *MongoClient) GetSuperChatByID(ctx context.Context, chatID int64) (*models.SuperChat, error) {
	var superChat models.SuperChat
	err := mc.Database.Collection(models.SuperChatCollectionName).FindOne(ctx, map[string]interface{}{"chat_id": chatID}).Decode(&superChat)
	if err != nil {
		return nil, err
	}

	return &superChat, nil
}

func (mc *MongoClient) UpdateSuperChat(ctx context.Context, superChat *models.SuperChat) error {
	_, err := mc.Database.Collection(models.SuperChatCollectionName).UpdateOne(ctx, map[string]interface{}{"chat_id": superChat.ChatID}, map[string]interface{}{"$set": superChat})
	return err
}
