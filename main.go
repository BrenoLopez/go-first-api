package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BrenoLopez/go-first-api/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.InitRoutes(&router.RouterGroup)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
