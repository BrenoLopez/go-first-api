package repository

import (
	"context"
	"os"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (userRepository *userRepository) Delete(id string) *httpError.HttpError {
	logger.Info("Init delete user repository", zap.String("jorney", "deleteUser"))
	collectionName := os.Getenv(MONGODB_COLLECTION)
	collection := userRepository.database.Collection(collectionName)

	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("jorney", "deleteUser"))
		return httpError.NewInternalServerError(err.Error())
	}
	return nil
}
