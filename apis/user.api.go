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

	c.IndentedJSON(http.StatusOK, common.SuccessAsData(services.GetUsers(paging)))
}

func GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	resp := services.FindUserById(userId)
	if errors.IsSuccess(resp.Code) {
		c.IndentedJSON(http.StatusOK, resp)
	} else {
		c.IndentedJSON(http.StatusOK, common.Failed(resp.Code))
	}
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	resp := services.CreateUser(user)
	common.UpsertResponse(c, resp)
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	resp := services.UpdateUser(user)
	common.UpsertResponse(c, resp)
}

func DeleteById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.Failed(errors.InvalidData))
		return
	}

	resp := services.DeleteUserById(userId)
	common.DeleteResponse(c, resp)
}
