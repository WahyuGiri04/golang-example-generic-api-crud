package main

import (
	"golang-example-generic-api-crud/config"
	"golang-example-generic-api-crud/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	port := os.Getenv("PORT")
	routes.SetupRoutes(r)

	r.Run(":" + port)
}
