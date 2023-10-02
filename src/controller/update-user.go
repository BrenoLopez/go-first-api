package controller

import (
	"net/http"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/BrenoLopez/go-first-api/src/config/validation"
	"github.com/BrenoLopez/go-first-api/src/controller/dto/request"
	"github.com/BrenoLopez/go-first-api/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (userController *userController) UpdateUser(context *gin.Context) {
	zapFields := zap.String("jorney", "updateUser")
	logger.Info("Init create user controller", zapFields)
	var userRequest request.UserUpdateRequestDto

	id := context.Param("id")

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info, error=%s", err)
		restError := validation.ValidateUserError(err)
		context.JSON(restError.Code, restError)
		return
	}
	userModel := model.NewUserUpdateModel(userRequest.Name, userRequest.Age)

	err := userController.service.Update(id, userModel)

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		restError := httpError.NewBadRequestError("Invalid userId, must be a hex value")
		context.JSON(restError.Code, restError)
	}

	if err != nil {
		context.JSON(err.Code, err)
	}

	logger.Info("User updated successfully", zapFields)

	context.Status(http.StatusOK)
}
