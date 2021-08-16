package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {
	var user models.User
	config.DB.Raw("SELECT * FROM users").Scan(&user)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := config.DB.Create(user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		// c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "failed to save"})
	} else {
		c.JSON(http.StatusOK, user)
	}

	c.JSON(http.StatusOK, gin.H{"message": "show detail"})
}

func ShowUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ShowUser"})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateUser"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteUser"})
}
