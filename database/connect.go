package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DatabaseConnect() (db *gorm.DB) {
	mysqlUser := "ngdangkiet"
	mysqlPass := "pwd"
	mysqlHost := "127.0.0.1"
	mysqlPort := "3307"
	mysqlDb := "user_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to db:", db)
	return db
}
