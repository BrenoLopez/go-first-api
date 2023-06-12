package view

import (
	"github.com/BrenoLopez/go-first-api/src/controller/dto/response"
	"github.com/BrenoLopez/go-first-api/src/model"
)

func CovertDomainToResponse(userModel model.UserModelInterface) response.UserResponseDto {
	return response.UserResponseDto{
		ID:    "",
		Email: userModel.GetEmail(),
		Name:  userModel.GetName(),
		Age:   userModel.GetAge(),
	}
}
