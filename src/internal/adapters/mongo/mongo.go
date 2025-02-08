package mongo

import (
	"context"
	"time"

	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	uri    string
	Client *mongo.Client
}

func (m *MongoClient) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(m.uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	m.Client = client

	return nil
}

func (m *MongoClient) Disconnect() error {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func NewMongoClient(uri string) (*MongoClient, error) {
	return &MongoClient{uri: uri}, nil
}
