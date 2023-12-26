package database

import (
	"fmt"
	"golang-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var DbError error

func Connect() {
	mysqlUser := "ngdangkiet"
	mysqlPass := "pwd"
	mysqlHost := "127.0.0.1"
	mysqlPort := "3307"
	mysqlDb := "user_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
	Instance, DbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DbError != nil {
		log.Fatal("Cannot connect to MySQL:", DbError)
	}

	log.Println("Connected to db:", Instance)
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Cannot migration:", err)
		return
	}
	log.Println("Database migration completed")
}
