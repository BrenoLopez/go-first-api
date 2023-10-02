package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"go.uber.org/zap"
)

func (userService *userService) Delete(id string) *httpError.HttpError {
	logger.Info("Init delete user service", zap.String("jorney", "updateUser"))

	err := userService.userRepository.Delete(id)

	if err != nil {
		return err
	}
	return nil
}
