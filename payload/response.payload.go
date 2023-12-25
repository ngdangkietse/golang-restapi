package payload

import "golang-rest-api/models"

type UserResponse struct {
	Code int          `json:"code"`
	Data *models.User `json:"data"`
}
