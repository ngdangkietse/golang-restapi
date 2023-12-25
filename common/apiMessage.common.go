package common

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api/errors"
	"net/http"
)

type ApiMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success() ApiMessage {
	return ApiMessage{Code: errors.Success, Message: "Success"}
}

func SuccessAsData(data interface{}) ApiMessage {
	return ApiMessage{Code: errors.Success, Message: "Success", Data: data}
}

func Failed(code int) ApiMessage {
	switch code {
	case errors.InvalidData:
		return ApiMessage{Code: errors.InvalidData, Message: "Invalid Data"}
	case errors.NotFound:
		return ApiMessage{Code: errors.NotFound, Message: "Not Found"}
	case errors.AlreadyExists:
		return ApiMessage{Code: errors.AlreadyExists, Message: "Already Exists"}
	default:
		return ApiMessage{Code: errors.Failed, Message: "Failed"}
	}
}

func UpsertResponse(c *gin.Context, respCode int) {
	if errors.IsSuccess(respCode) {
		c.IndentedJSON(http.StatusOK, SuccessAsData(respCode))
	} else {
		c.IndentedJSON(http.StatusOK, Failed(respCode))
	}
}

func DeleteResponse(c *gin.Context, respCode int) {
	if errors.IsSuccess(respCode) {
		c.IndentedJSON(http.StatusOK, Success())
	} else {
		c.IndentedJSON(http.StatusOK, Failed(respCode))
	}
}
