package controllers

import (
	"base-project-go/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.Login
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"user": user.Username, "pass": user.Password})
}

func Register(c *gin.Context) {
	var user models.Login

	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	// c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, "register")
}
