package controller

import (
	"net/http"
	"net/mail"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (userController *userController) FindUserById(context *gin.Context) {
	id := context.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errorMessage := httpError.NewBadRequestError("UserId is not a valid id")

		context.JSON(errorMessage.Code, errorMessage)
		return
	}

	userModel, err := userController.service.FindUserById(id)

	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.CovertDomainToResponse(userModel))
}

func (userController *userController) FindUserByEmail(context *gin.Context) {
	email := context.Param("email")

	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := httpError.NewBadRequestError("Email is not valid")

		context.JSON(errorMessage.Code, errorMessage)
		return
	}

	userModel, err := userController.service.FindUserByEmail(email)

	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.JSON(http.StatusOK, view.CovertDomainToResponse(userModel))
}
