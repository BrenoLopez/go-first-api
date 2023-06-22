package controller

import (
	"net/http"

	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/config/validation"
	"github.com/BrenoLopez/go-first-api/src/controller/dto/request"
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/BrenoLopez/go-first-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (userController *userController) CreateUser(context *gin.Context) {
	zapFields := zap.String("jorney", "createUser")
	logger.Info("Init create user controller", zapFields)
	var userRequest request.UserRequestDto

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info, error=%s", err)
		restError := validation.ValidateUserError(err)
		context.JSON(restError.Code, restError)
		return
	}
	userModel := model.NewUserModel(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := userController.service.Create(userModel); err != nil {
		context.JSON(err.Code, err)
	}

	logger.Info("User created successfully", zapFields)
	context.JSON(http.StatusOK, view.CovertDomainToResponse(userModel))
}
