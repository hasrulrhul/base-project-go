package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@(127.0.0.1:8889)/" + os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	config := gorm.Config{}

	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		panic("Failed to connect to database!" + os.Getenv("DB_USER"))
	}
	DB = db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

// //SetupDatabaseConnection is creating a new connection to our database
// func SetupDatabaseConnection() {
// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 		panic("Failed to load env file")
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbName := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to create a connection to database")
// 	}
// 	//nanti kita isi modelnya di sini
// 	// db.AutoMigrate(&entity.Book{}, &entity.User{})
// 	return db
// }

// //CloseDatabaseConnection method is closing a connection between your app and your db
// func CloseDatabaseConnection(db *gorm.DB) {
// 	dbSQL, err := db.DB()
// 	if err != nil {
// 		panic("Failed to close connection from database")
// 	}
// 	dbSQL.Close()
// }
