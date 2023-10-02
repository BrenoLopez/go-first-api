package controller

import (
	"net/http"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/BrenoLopez/go-first-api/src/config/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (userController *userController) DeleteUser(context *gin.Context) {
	zapFields := zap.String("jorney", "deleteUser")
	logger.Info("Init create user controller", zapFields)

	id := context.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		restError := httpError.NewBadRequestError("Invalid userId, must be a hex value")
		context.JSON(restError.Code, restError)
	}

	err := userController.service.Delete(id)

	if err != nil {
		context.JSON(err.Code, err)
		return
	}
	context.Status(http.StatusOK)
}
