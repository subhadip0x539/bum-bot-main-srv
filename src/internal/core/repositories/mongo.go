package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *MongoRepoImpl) FindOne(collection string, filter interface{}, result interface{}) (bool, error) {
	c := r.client.Database(r.database).Collection(collection)

	if err := c.FindOne(context.TODO(), filter).Decode(result); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *MongoRepoImpl) InsertMany(collection string, documents []interface{}) error {
	c := r.client.Database(r.database).Collection(collection)

	if _, err := c.InsertMany(context.TODO(), documents); err != nil {
		return err
	}

	return nil
}

func (r *MongoRepoImpl) Aggregate(collection string, pipeline []bson.M, results interface{}) error {
	c := r.client.Database(r.database).Collection(collection)

	cursor, err := c.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return err

	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), results); err != nil {
		return err
	}

	return nil
}

func NewMongoRepo(client *mongo.Client, database string) *MongoRepoImpl {
	return &MongoRepoImpl{client: client, database: database}
}
