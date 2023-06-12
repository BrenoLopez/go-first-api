package controller

import (
	"github.com/BrenoLopez/go-first-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &userController{service}
}

type UserControllerInterface interface {
	CreateUser(context *gin.Context)
	DeleteUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	FindUserById(context *gin.Context)
	FindUserByEmail(context *gin.Context)
}

type userController struct {
	service service.UserServiceInterface
}
