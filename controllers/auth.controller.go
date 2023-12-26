package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api/models"
	"golang-rest-api/payload"
	"golang-rest-api/services"
	"net/http"
)

func Authenticate(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, payload.HandleException(err.Error()))
		return
	}

	_, data := services.Authenticate(loginRequest)

	c.IndentedJSON(http.StatusOK, data)
}
