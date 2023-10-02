package repository

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database: database}
}

type userRepository struct {
	database *mongo.Database
}

type UserRepository interface {
	Create(userModel model.UserModelInterface) (model.UserModelInterface, *httpError.HttpError)
	FindUserById(id string) (model.UserModelInterface, *httpError.HttpError)
	FindUserByEmail(email string) (model.UserModelInterface, *httpError.HttpError)
	Update(id string, userModel model.UserModelInterface) *httpError.HttpError
	Delete(id string) *httpError.HttpError
}
