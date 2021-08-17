package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexMenu(c *gin.Context) {
	var menu []models.Menu
	config.DB.Find(&menu)
	c.JSON(http.StatusOK, menu)
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
		panic(err)
	}
	// config.DB.Create(&menu)
	if err := config.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func ShowMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var menu models.Menu
	err := config.DB.First(&menu, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
	} else {
		c.JSON(http.StatusOK, menu)
	}
}

func UpdateMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var menu models.Menu
	err := config.DB.First(&menu, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := c.BindJSON(&menu); err != nil {
		panic(err)
	}
	if err := config.DB.Updates(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func DeleteMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var menu models.Menu
	err := config.DB.First(&menu, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
