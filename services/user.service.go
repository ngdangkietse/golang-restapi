package services

import (
	"golang-rest-api/database"
	"golang-rest-api/models"
	"golang-rest-api/payload"
	"golang-rest-api/utils"
	"golang-rest-api/validators"
)

//var logger = shared.ConfigLogger()

func CreateUser(user models.User) (bool, interface{}) {
	db := database.Instance

	errorMap := validators.ValidatorUser(user, true)
	if len(errorMap) > 0 {
		return false, payload.HandleInvalidData(errorMap)
	}

	success, _ := FindUserByEmail(user.Email)
	if success {
		return false, payload.HandleAlreadyExists("User Email", user.Email)
	}

	if err := user.HashPassword(user.Password); err != nil {
		return false, payload.HandleException(err.Error())
	}

	if err := db.Table(user.TableName()).
		Create(&user).Error; err != nil {
		return false, payload.HandleException(err.Error())
	}

	return true, payload.HandleSuccess(user.Id)
}

func FindUserByEmail(email string) (bool, interface{}) {
	db := database.Instance

	var user models.User
	record := db.Table(user.TableName()).Where("email = ?", email).First(&user)

	if record.Error != nil {
		return false, payload.HandleNotFound("Email", email)
	}

	return true, payload.HandleSuccess(user)
}

func FindUserById(userId int) (bool, interface{}) {
	db := database.Instance
	var user models.User
	if err := db.Table(user.TableName()).
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		return false, payload.HandleNotFound("UserId", userId)
	}

	return true, payload.HandleSuccess(user)
}

func UpdateUser(user models.User) (bool, interface{}) {
	db := database.Instance

	errorMap := validators.ValidatorUser(user, false)
	if len(errorMap) > 0 {
		return false, payload.HandleInvalidData(errorMap)
	}

	success, _ := FindUserById(user.Id)

	if !success {
		return false, payload.HandleNotFound("UserId", user.Id)
	}

	if err := user.HashPassword(user.Password); err != nil {
		return false, payload.HandleException(err.Error())
	}

	if err := db.Table(user.TableName()).
		Updates(&user).Error; err != nil {
		return false, payload.HandleException(err.Error())
	}

	return true, payload.HandleSuccess(user.Id)
}

func DeleteUserById(userId int) (bool, interface{}) {
	db := database.Instance

	success, _ := FindUserById(userId)

	if !success {
		return false, payload.HandleNotFound("UserId", userId)
	}

	if err := db.Table("tbl_user").
		Where("id = ?", userId).
		Delete(nil).Error; err != nil {
		return false, payload.HandleException(err.Error())
	}
	return true, payload.HandleSuccess(nil)
}

func GetUsers(paging models.Paging) (bool, interface{}) {
	db := database.Instance
	var users []models.User

	paging = utils.Paging(paging)

	var user models.User
	if err := db.Table(user.TableName()).
		Limit(paging.PageSize).
		Offset(utils.PageOffset(paging.PageIndex, paging.PageSize)).
		Order(paging.PageSort + " " + paging.PageDirection).
		Find(&users).Error; err != nil {
		return true, payload.HandleSuccess([]models.User{})
	}

	return true, payload.HandleSuccess(users)
}
