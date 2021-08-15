package main

import (
	"base-project-go/database"
	"base-project-go/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Menu{})
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Option{})
	database.DB.AutoMigrate(&models.Role{})
	database.DB.AutoMigrate(&models.UserMenu{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Gin",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		var result models.User
		database.DB.Raw("SELECT * FROM users").Scan(&result)
		c.JSON(http.StatusOK, result)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		var result models.User
		ID := c.Param("id")
		database.DB.Raw("SELECT * FROM users WHERE id = ?", ID).Scan(&result)
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": result})
	})

	r.Run(":" + os.Getenv("APP_PORT"))
}
