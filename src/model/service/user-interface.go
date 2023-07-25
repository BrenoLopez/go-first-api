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
	Update(string, model.UserModelInterface) *httpError.HttpError
	Find(string) (*model.UserModelInterface, *httpError.HttpError)
	Delete(string) *httpError.HttpError
}
