package repository

import (
	"context"
	"os"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (userRepository *userRepository) Update(id string, userModel model.UserModelInterface) *httpError.HttpError {
	logger.Info("Init update user repository", zap.String("jorney", "updateUser"))
	collectionName := os.Getenv(MONGODB_COLLECTION)
	collection := userRepository.database.Collection(collectionName)
	value := converter.ConvertDomainToEntity(userModel)
	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	updateFields := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, updateFields)

	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("jorney", "updateUser"))
		return httpError.NewInternalServerError(err.Error())
	}
	return nil
}
