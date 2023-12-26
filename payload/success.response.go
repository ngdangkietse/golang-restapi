package payload

import "golang-rest-api/errors"

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func HandleSuccess(data interface{}) SuccessResponse {
	return SuccessResponse{
		Code: errors.Success,
		Data: data,
	}
}
