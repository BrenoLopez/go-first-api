package routes

import (
	"github.com/BrenoLopez/go-first-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup) {
	route.GET("/users/:id", controller.FindUserById)
	route.GET("/users-email/:email", controller.FindUserByEmail)
	route.POST("/users", controller.CreateUser)
	route.PUT("/users/:id")
	route.DELETE("/users/:id")
}
