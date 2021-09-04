package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUserMenu(c *gin.Context) {
	var usermenu []models.UserMenu
	config.DB.Preload("Role").Preload("Menu").Find(&usermenu)
	c.JSON(http.StatusOK, service.Response(usermenu, c, "", 0))
}

func CreateUserMenu(c *gin.Context) {
	var usermenu models.UserMenu
	if err := c.BindJSON(&usermenu); err != nil {
		panic(err)
	}
	if err := config.DB.Create(&usermenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func ShowUserMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var usermenu models.UserMenu
	err := config.DB.Preload("Role").Preload("Menu").First(&usermenu, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
	} else {
		c.JSON(http.StatusOK, usermenu)
	}
}

func UpdateUserMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var usermenu models.UserMenu
	err := config.DB.First(&usermenu, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := c.BindJSON(&usermenu); err != nil {
		panic(err)
	}
	if err := config.DB.Updates(&usermenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func DeleteUserMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var usermenu models.UserMenu
	err := config.DB.First(&usermenu, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&usermenu).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
