package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepoImpl struct {
	client   *mongo.Client
	database string
}

func (r *MongoRepoImpl) InsertOne(collection string, document interface{}) error {
	c := r.client.Database(r.database).Collection(collection)

	if _, err := c.InsertOne(context.TODO(), document); err != nil {
		return err
	}

	return nil
}

func (r *MongoRepoImpl) InsertMany(collection string, documents []interface{}) error {
	c := r.client.Database(r.database).Collection(collection)

	if _, err := c.InsertMany(context.TODO(), documents); err != nil {
		return err
	}

	return nil
}

func NewMongoRepo(client *mongo.Client, database string) *MongoRepoImpl {
	return &MongoRepoImpl{client: client, database: database}
}
