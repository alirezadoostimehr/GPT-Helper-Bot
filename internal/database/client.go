package database

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"
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
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(config.GlobalConfig.Mongo.URI),
		options.Client().SetAuth(options.Credential{
			Username: config.GlobalConfig.Mongo.USERNAME,
			Password: config.GlobalConfig.Mongo.PASSWORD,
		}),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
