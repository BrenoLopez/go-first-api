package repository

import (
	"context"
	"fmt"
	"os"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/model/repository/entity"
	"github.com/BrenoLopez/go-first-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (userRepository *userRepository) FindUserById(id string) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init find user by id repository", zap.String("jorney", "FindUserById"))

	collectionName := os.Getenv(MONGODB_COLLECTION)

	collection := userRepository.database.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		var errorMessage string
		if err == mongo.ErrNoDocuments {
			errorMessage = fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("jorney", "findUserById"))
			return nil, httpError.NewBadRequestError(errorMessage)
		}
		errorMessage = "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("jorney", "FindUserById"))
		return nil, httpError.NewInternalServerError(errorMessage)
	}
	logger.Info("Finish find user by id repository", zap.String("jorney", "findUserById"), zap.String("userId", userEntity.Id.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (userRepository *userRepository) FindUserByEmail(email string) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init find user by email repository", zap.String("jorney", "FindUserByEmail"))

	collectionName := os.Getenv(MONGODB_COLLECTION)

	collection := userRepository.database.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		var errorMessage string
		if err == mongo.ErrNoDocuments {
			errorMessage = fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("jorney", "FindUserByEmail"))
			return nil, httpError.NewBadRequestError(errorMessage)
		}
		errorMessage = "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("jorney", "FindUserByEmail"))
		return nil, httpError.NewInternalServerError(errorMessage)
	}
	logger.Info("Finish find user by email repository", zap.String("jorney", "FindUserByEmail"),
		zap.String("email", email), zap.String("userId", userEntity.Id.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}
