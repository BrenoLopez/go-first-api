package main

import (
	"github.com/BrenoLopez/go-first-api/src/controller"
	"github.com/BrenoLopez/go-first-api/src/model/repository"
	"github.com/BrenoLopez/go-first-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repository := repository.NewUserRepository(database)
	service := service.NewUserService(repository)
	return controller.NewUserController(service)
}
