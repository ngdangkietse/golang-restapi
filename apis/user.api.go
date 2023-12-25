package apis

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api/common"
	"golang-rest-api/errors"
	"golang-rest-api/models"
	"golang-rest-api/services"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	var paging models.Paging

	if err := c.BindJSON(&paging); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	c.IndentedJSON(http.StatusOK, services.GetUsers(paging))
}

func GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	var resp = services.FindUserById(userId)
	if errors.IsSuccess(resp.Code) {
		c.IndentedJSON(http.StatusOK, resp)
	} else {
		c.IndentedJSON(http.StatusOK, common.Failed(resp.Code))
	}
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	resp := services.CreateUser(user)
	if errors.IsSuccess(resp) {
		c.IndentedJSON(http.StatusOK, resp)
	} else {
		c.IndentedJSON(http.StatusOK, common.Failed(resp))
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	var resp = services.UpdateUser(user)
	if errors.IsSuccess(resp) {
		c.IndentedJSON(http.StatusOK, resp)
	} else {
		c.IndentedJSON(http.StatusOK, common.Failed(resp))
	}
}

func DeleteById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	var resp = services.DeleteUserById(userId)
	if errors.IsSuccess(resp) {
		c.IndentedJSON(http.StatusOK, common.Success())
	} else {
		c.IndentedJSON(http.StatusOK, common.Failed(resp))
	}
}
