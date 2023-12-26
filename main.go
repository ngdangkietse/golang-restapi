package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-rest-api/controllers"
	"golang-rest-api/database"
	"golang-rest-api/middlewares"
	"log"
)

func main() {
	// env
	loadEnv()

	// Database
	loadDatabase()

	// Server
	serverApplication()
}

func loadDatabase() {
	database.Connect()
	database.Migrate()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env:", err)
	}
	log.Println(".env file loaded successfully")
}

func serverApplication() {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/login", controllers.Authenticate)
		api.POST("/users", controllers.CreateUser)

		secured := api.Group("/secured").Use(middlewares.Authentication())
		{
			secured.GET("/users", controllers.GetUsers)
			secured.GET("/users/detail/:id", controllers.GetUserById)
			secured.PUT("/users", controllers.UpdateUser)
			secured.DELETE("/users/:id", controllers.DeleteById)
		}
	}

	err := router.Run("localhost:6969")
	if err != nil {
		log.Fatal("Error router:", err)
		return
	}
	log.Println("Server running on port 6969")
}
