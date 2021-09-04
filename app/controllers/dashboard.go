package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "welcome to dashboard"})
}

func IndexPost(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	config.DB.Create(&role)
	c.JSON(200, gin.H{"msg": role}) // Your custom response here
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello"}) // Your custom response here
}
