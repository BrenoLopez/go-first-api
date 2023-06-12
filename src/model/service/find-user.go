package service

import (
	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/model"
)

func (*userService) Find(id string) (*model.UserModelInterface, *httpError.HttpError) {
	return nil, nil
}
