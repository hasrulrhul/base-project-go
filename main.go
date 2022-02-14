package main

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/route"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Menu{})
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Option{})
	config.DB.AutoMigrate(&models.Role{})
	config.DB.AutoMigrate(&models.UserMenu{})
}

func main() {
	defer config.CloseDatabaseConnection(config.DB)

	r := route.SetupRouter()
	r.Run(":" + os.Getenv("APP_PORT"))
}
