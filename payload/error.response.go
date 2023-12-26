package payload

import (
	"fmt"
	"golang-rest-api/errors"
)

type ErrorResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func HandleException(message string) ErrorResponse {
	return ErrorResponse{
		Code:    errors.Failed,
		Message: message,
	}
}

func HandleInvalidData(errorMap map[string]string) ErrorResponse {
	return ErrorResponse{
		Code:    errors.InvalidData,
		Message: "",
		Errors:  errorMap,
	}
}

func HandleNotFound(key interface{}, value interface{}) ErrorResponse {
	return ErrorResponse{
		Code:    errors.NotFound,
		Message: fmt.Sprintf("Not found [%v] with value [%v]", key, value),
	}
}

func HandleAlreadyExists(key interface{}, value interface{}) ErrorResponse {
	return ErrorResponse{
		Code:    errors.AlreadyExists,
		Message: fmt.Sprintf("[%v] already exists with value [%v]", key, value),
	}
}
