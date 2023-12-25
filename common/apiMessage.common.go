package common

import "golang-rest-api/errors"

type ApiMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success() ApiMessage {
	return ApiMessage{errors.Success, "Success"}
}

func Failed(code int) ApiMessage {
	switch code {
	case errors.InvalidData:
		return ApiMessage{errors.InvalidData, "Invalid Data"}
	case errors.NotFound:
		return ApiMessage{errors.NotFound, "Not Found"}
	case errors.AlreadyExists:
		return ApiMessage{errors.AlreadyExists, "Already Exists"}
	default:
		return ApiMessage{errors.Failed, "Failed"}
	}
}
