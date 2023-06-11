package controller

import (
	"fmt"

	"github.com/BrenoLopez/go-first-api/src/config/validation"
	"github.com/BrenoLopez/go-first-api/src/controller/dto/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequestDto

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		context.JSON(restError.Code, restError)
		return
	}
	fmt.Println(userRequest)
}
