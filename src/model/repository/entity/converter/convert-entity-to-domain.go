package converter

import (
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserModelInterface {
	domain := model.NewUserModel(entity.Email, entity.Password, entity.Name, entity.Age)

	domain.SetId(entity.Id.Hex())

	return domain
}
