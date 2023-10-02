package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/model/repository"
)

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

type userService struct {
	userRepository repository.UserRepository
}

type UserServiceInterface interface {
	Create(model.UserModelInterface) (model.UserModelInterface, *httpError.HttpError)
	Update(id string, userModel model.UserModelInterface) *httpError.HttpError
	FindUserById(id string) (model.UserModelInterface, *httpError.HttpError)
	FindUserByEmail(email string) (model.UserModelInterface, *httpError.HttpError)
	Delete(string) *httpError.HttpError
}
