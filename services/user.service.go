package services

import (
	"fmt"
	"golang-rest-api/database"
	"golang-rest-api/errors"
	"golang-rest-api/models"
	"golang-rest-api/payload"
	"golang-rest-api/shared"
	"golang-rest-api/utils"
	"log"
)

var logger = shared.ConfigLogger()

func CreateUser(user models.User) int {
	db := database.Connect()
	if user.Name == "" || user.Address == "" {
		return errors.InvalidData
	}

	if err := db.Table(models.User{}.TableName()).
		Create(&user).Error; err != nil {
		logger.Error("Error when creating user!")
		return errors.Failed
	}

	return user.Id
}

func FindUserById(userId int) *payload.UserResponse {
	db := database.Connect()
	var user models.User
	if err := db.Table(models.User{}.TableName()).
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		logger.Error(fmt.Sprintf("User ID [%d] not found!", userId))
		return &payload.UserResponse{
			Code: errors.NotFound,
		}
	}

	return &payload.UserResponse{
		Code: errors.Success,
		Data: &user,
	}
}

func UpdateUser(user models.User) int {
	if user.Name == "" || user.Address == "" {
		return errors.InvalidData
	}

	if FindUserById(user.Id).Code == errors.NotFound {
		return errors.NotFound
	}

	db := database.Connect()
	if err := db.Table(models.User{}.TableName()).
		Updates(&user).Error; err != nil {
		logger.Error("Error when updating user!")
		return errors.Failed
	}

	return user.Id
}

func DeleteUserById(userId int) int {
	if FindUserById(userId).Code == errors.NotFound {
		return errors.NotFound
	}
	db := database.Connect()
	if err := db.Table(models.User{}.TableName()).
		Where("id = ?", userId).
		Delete(nil).Error; err != nil {
		log.Fatal("Error when deleting user:", err)
		return errors.Failed
	}
	return errors.Success
}

func GetUsers(paging models.Paging) []models.User {
	db := database.Connect()
	var users []models.User

	paging = utils.Paging(paging)

	if err := db.Table(models.User{}.TableName()).
		Limit(paging.PageSize).
		Offset(utils.PageOffset(paging.PageIndex, paging.PageSize)).
		Order(paging.PageSort + " " + paging.PageDirection).
		Find(&users).Error; err != nil {
		return []models.User{}
	}

	return users
}
