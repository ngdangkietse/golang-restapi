package main

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api/apis"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/users", apis.GetUsers)
	router.GET("/api/v1/users/detail/:id", apis.GetUserById)
	router.POST("/api/v1/users", apis.CreateUser)
	router.PUT("/api/v1/users", apis.UpdateUser)
	router.DELETE("/api/v1/users/:id", apis.DeleteById)

	err := router.Run("localhost:6969")
	if err != nil {
		return
	}
}
