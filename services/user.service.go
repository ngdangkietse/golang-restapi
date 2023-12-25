package services

import (
	"golang-rest-api/database"
	"golang-rest-api/errors"
	"golang-rest-api/models"
	"golang-rest-api/payload"
	"log"
)

func CreateUser(user models.User) int {
	db := database.DatabaseConnect()
	if user.Name == "" || user.Address == "" {
		return errors.InvalidData
	}

	if err := db.Table(models.User{}.TableName()).
		Create(&user).Error; err != nil {
		log.Fatal("Error when creating user:", err)
		return errors.Failed
	}

	return user.Id
}

func FindUserById(userId int) *payload.UserResponse {
	db := database.DatabaseConnect()
	var user models.User
	if err := db.Table(models.User{}.TableName()).
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		log.Fatal("Error when getting user:", err)
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

	db := database.DatabaseConnect()
	if err := db.Table(models.User{}.TableName()).
		Where("id = ?", user.Id).
		Updates(&user).Error; err != nil {
		log.Fatal("Error when updating user:", err)
		return errors.Failed
	}

	return user.Id
}

func DeleteUserById(userId int) int {
	if FindUserById(userId).Code == errors.NotFound {
		return errors.NotFound
	}
	db := database.DatabaseConnect()
	if err := db.Table(models.User{}.TableName()).
		Where("id = ?", userId).
		Delete(nil).Error; err != nil {
		log.Fatal("Error when deleting user:", err)
		return errors.Failed
	}
	return errors.Success
}

func GetUsers(paging models.Paging) []models.User {
	db := database.DatabaseConnect()
	var users []models.User

	if paging.PageSize <= 0 {
		paging.PageSize = 1
	}

	if paging.PageLimit <= 0 {
		paging.PageLimit = 10
	}

	offset := (paging.PageSize - 1) * paging.PageLimit

	if err := db.Table(models.User{}.TableName()).
		Limit(paging.PageLimit).
		Offset(offset).
		Find(&users).Error; err != nil {
		log.Fatal("Error when deleting user:", err)
		return []models.User{}
	}

	return users
}
