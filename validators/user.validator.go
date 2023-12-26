package validators

import (
	"golang-rest-api/constants"
	"golang-rest-api/models"
)

func ValidatorUser(user models.User, isCreate bool) map[string]string {
	errorMap := make(map[string]string)

	if !isCreate && user.Id == 0 {
		errorMap["UserId"] = constants.MissingData
	}
	if user.Name == "" {
		errorMap["Name"] = constants.MissingData
	}
	if user.Address == "" {
		errorMap["Address"] = constants.MissingData
	}
	if user.Email == "" {
		errorMap["Email"] = constants.MissingData
	}
	if user.Password == "" {
		errorMap["Password"] = constants.MissingData
	}

	return errorMap
}
