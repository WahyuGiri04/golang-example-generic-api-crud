package main

import (
	"golang-example-generic-api-crud/config"
	"golang-example-generic-api-crud/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDatabase()

	// Inisialisasi Gin
	r := gin.Default()

	// Menentukan port service
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8080" // Default port jika tidak ada di env
	}

	// Registrasi service ke Consul
	config.RegisterServiceWithConsul()

	// Setup routes
	routes.SetupRoutes(r)

	// Endpoint health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Jalankan server
	log.Printf("Server berjalan di port %s...", port)
	r.Run(":" + port)
}
