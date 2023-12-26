package services

import (
	"fmt"
	"golang-rest-api/auth"
	"golang-rest-api/models"
	"golang-rest-api/payload"
)

func Authenticate(request models.LoginRequest) (bool, interface{}) {
	// TODO: validate email and password
	success, data := FindUserByEmail(request.Email)
	if !success {
		return false, payload.HandleNotFound("Email", request.Email)
	}

	successResponse, okResponse := data.(payload.SuccessResponse)
	if !okResponse {
		return false, payload.HandleException("Conversion failed!")
	}

	user, okUser := successResponse.Data.(models.User)
	if !okUser {
		return false, payload.HandleException("Conversion failed!")
	}

	// TODO: generate token
	credentialErr := user.CheckMatchesPassword(request.Password)
	if credentialErr != nil {
		return false, payload.HandleException(fmt.Sprintf("Invalid password. Error: [%v]", credentialErr.Error()))
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return false, payload.HandleException(err.Error())
	}

	return true, payload.HandleSuccess(models.LoginResponse{
		Token:  token,
		UserId: user.Id,
		RoleId: user.RoleId,
	})
}
