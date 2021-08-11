package main

import (
	"base-project-go/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDatabase()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Gin",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		var result Users
		database.DB.Raw("SELECT * FROM users").Scan(&result)
		c.JSON(http.StatusOK, result)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		var result Users
		ID := c.Param("id")
		database.DB.Raw("SELECT * FROM users WHERE id = ?", ID).Scan(&result)
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": result})
	})

	r.Run(":" + os.Getenv("APP_PORT"))
}
