package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
	"go.uber.org/zap"
)

func (userService *userService) FindUserById(id string) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init find user by id service", zap.String("jorney", "FindUserById"))
	return userService.userRepository.FindUserById(id)
}

func (userService *userService) FindUserByEmail(email string) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init find user by email service", zap.String("jorney", "FindUserByEmail"))
	return userService.userRepository.FindUserByEmail(email)
}
