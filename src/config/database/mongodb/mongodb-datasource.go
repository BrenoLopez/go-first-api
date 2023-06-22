package mongodb

import (
	"context"
	"os"

	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL  = "MONGODB_URL"
	MONGODB_NAME = "MONGODB_NAME"
)

func NewMongoDbConnection(context context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(os.Getenv(MONGODB_URL)))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context, nil); err != nil {
		return nil, err
	}
	logger.Info("Mongodb connection successfully")
	return client.Database(os.Getenv(MONGODB_NAME)), nil
}
