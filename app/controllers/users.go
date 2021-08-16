package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {
	var user []models.User
	config.DB.Preload("Role").Find(&user)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// func CreateOrUpdateUser(c *gin.Context) {
// 	var user []models.User
// 	if err := c.BindJSON(&user); err != nil {
// 		panic(err)
// 	}
// 	config.DB.Clauses(clause.OnConflict{
// 		Columns:   []clause.Column{{Name: "id"}},
// 		DoUpdates: clause.AssignmentColumns([]string{"name", "username", "email", "password", "role_id"}),
// 	}).Create(&user)
// 	c.JSON(http.StatusOK, user)
// }

func ShowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	config.DB.Preload("Role").First(&user, id)
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.Preload("Role").First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Record not found")
	}
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	config.DB.Updates(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Record not found")
	}
	config.DB.Delete(&user)
}
