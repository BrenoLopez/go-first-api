package converter

import (
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/model/repository/entity"
)

func ConvertDomainToEntity(userModel model.UserModelInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    userModel.GetEmail(),
		Password: userModel.GetPassword(),
		Name:     userModel.GetName(),
		Age:      userModel.GetAge(),
	}
}
