package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
	"go.uber.org/zap"
)

func (userService *userService) Create(userModel model.UserModelInterface) (model.UserModelInterface, *httpError.HttpError) {
	logger.Info("Init create user model", zap.String("jorney", "createUser"))

	userModel.EncryptPassword()

	userDomain, err := userService.userRepository.Create(userModel)

	if err != nil {
		return nil, err
	}
	return userDomain, nil
}
