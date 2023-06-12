package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/model"
)

func NewUserService() *userService {
	return &userService{}
}

type userService struct {
}

type UserServiceInterface interface {
	Create(model.UserModelInterface) *httpError.HttpError
	Update(string, model.UserModelInterface) *httpError.HttpError
	Find(string) (*model.UserModelInterface, *httpError.HttpError)
	Delete(string) *httpError.HttpError
}
