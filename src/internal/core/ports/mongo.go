package ports

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepo interface {
	InsertOne(collation string, document interface{}) error
	InsertMany(collation string, document []interface{}) error
	FindOne(collation string, filter interface{}, result interface{}) (bool, error)
	Aggregate(collection string, pipeline []bson.M, results interface{}) error
}
