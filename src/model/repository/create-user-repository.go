package repository

import (
	"context"
	"os"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

func (userRepository *userRepository) Create(userModel model.UserModelInterface) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init create user repository")
	collectionName := os.Getenv(MONGODB_COLLECTION)
	collection := userRepository.database.Collection(collectionName)
	value := converter.ConvertDomainToEntity(userModel)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, httpError.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
