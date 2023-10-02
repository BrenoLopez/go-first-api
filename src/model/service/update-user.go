package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/model"
	"go.uber.org/zap"
)

func (userService *userService) Update(id string, userModel model.UserModelInterface) *httpError.HttpError {
	logger.Info("Init update user service", zap.String("jorney", "updateUser"))

	err := userService.userRepository.Update(id, userModel)

	if err != nil {
		return err
	}
	return nil
}
