package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@(127.0.0.1:3306)/" + os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	config := gorm.Config{}

	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		panic("Failed to connect to database!" + os.Getenv("DB_USER"))
	}
	DB = db

}
