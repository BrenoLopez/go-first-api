package repository

import (
	"context"
	"os"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
)

const (
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

func (userRepository *userRepository) Create(userModel model.UserModelInterface) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init create user repository")
	collectionName := os.Getenv(MONGODB_COLLECTION)
	collection := userRepository.database.Collection(collectionName)
	value, err := userModel.GetJSONValue()
	if err != nil {
		return nil, httpError.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, httpError.NewInternalServerError(err.Error())
	}

	userModel.SetId(result.InsertedID.(string))
	return userModel, nil
}
