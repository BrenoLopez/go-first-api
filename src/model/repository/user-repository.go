package repository

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database: database}
}

type userRepository struct {
	database *mongo.Database
}

type UserRepository interface {
	Create(userModel model.UserModelInterface) (model.UserModelInterface, *httpError.HttpError)
}
