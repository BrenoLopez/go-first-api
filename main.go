package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BrenoLopez/go-first-api/src/config/database/mongodb"
	"github.com/BrenoLopez/go-first-api/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	context := context.Background()

	database, err := mongodb.NewMongoDbConnection(context)

	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.InitRoutes(&router.RouterGroup, userController)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
