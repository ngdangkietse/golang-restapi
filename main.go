package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-rest-api/controllers"
	"golang-rest-api/database"
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

	router.GET("/api/v1/users", controllers.GetUsers)
	router.GET("/api/v1/users/detail/:id", controllers.GetUserById)
	router.POST("/api/v1/users", controllers.CreateUser)
	router.PUT("/api/v1/users", controllers.UpdateUser)
	router.DELETE("/api/v1/users/:id", controllers.DeleteById)

	err := router.Run("localhost:6969")
	if err != nil {
		log.Fatal("Error router:", err)
		return
	}
	log.Println("Server running on port 6969")
}
