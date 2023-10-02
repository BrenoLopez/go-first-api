package routes

import (
	"github.com/BrenoLopez/go-first-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup, userController controller.UserControllerInterface) {
	route.GET("/users/:id", userController.FindUserById)
	route.GET("/users-email/:email", userController.FindUserByEmail)
	route.POST("/users", userController.CreateUser)
	route.PUT("/users/:id", userController.UpdateUser)
	route.DELETE("/users/:id", userController.DeleteUser)
}
