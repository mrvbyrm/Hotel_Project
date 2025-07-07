package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"stajprojesi/config"
	"stajprojesi/routes"
)

func main() {
	// Connect to the database
	err := config.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Database connection successful")
	config.Setup()

	// Get the port from environment variable or default to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Create a new Gin router
	router := gin.New()
	router.Use(gin.Logger()) // Log all requests

	// Define route groups
	api := router.Group("/api")
	{
		routes.CustomerRoutes(api.Group("/customers"))
		routes.RoomTypeRoutes(api.Group("/roomTypes"))
		routes.RoomRoutes(api.Group("/rooms"))
		routes.PaymentRoutes(api.Group("/payments"))
		routes.UserRoutes(api.Group("/users"))
		routes.ReservationRoot(api.Group("/reservations"))
		routes.AdminRoutes(api.Group("/admins"))
	}

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
