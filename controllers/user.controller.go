package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api/models"
	"golang-rest-api/payload"
	"golang-rest-api/services"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	var paging models.Paging

	if err := c.BindJSON(&paging); err != nil {
		c.IndentedJSON(http.StatusBadRequest, payload.HandleException(err.Error()))
		return
	}

	success, data := services.GetUsers(paging)

	if success {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, payload.HandleException(err.Error()))
		return
	}

	_, data := services.FindUserById(userId)

	c.IndentedJSON(http.StatusOK, data)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, payload.HandleException(err.Error()))
		return
	}

	_, data := services.CreateUser(user)

	c.IndentedJSON(http.StatusOK, data)
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, payload.HandleException(err.Error()))
		return
	}

	_, data := services.UpdateUser(user)

	c.IndentedJSON(http.StatusOK, data)
}

func DeleteById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, payload.HandleException(err.Error()))
		return
	}

	_, data := services.DeleteUserById(userId)

	c.IndentedJSON(http.StatusOK, data)
}
