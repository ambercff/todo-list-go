package database

import (
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB * gorm.DB

func Connect() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
}