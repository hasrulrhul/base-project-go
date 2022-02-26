package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexMenu(c *gin.Context) {
	var menu []models.Menu
	config.DB.Find(&menu)
	response := helper.BuildResponse(true, "List of menu!", menu)
	c.JSON(http.StatusOK, response)
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	if err := config.DB.Create(&menu).Error; err != nil {
		response := helper.BuildErrorResponse("Created menu failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Created menu successfull!", menu)
		c.JSON(http.StatusCreated, response)
	}
}

func ShowMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var menu models.Menu
	err := config.DB.First(&menu, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Menu not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Detail menu!", menu)
		c.JSON(http.StatusCreated, response)
	}
}

func UpdateMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var menu models.Menu
	err := config.DB.First(&menu, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Menu not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := c.BindJSON(&menu); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	if err := config.DB.Updates(&menu).Error; err != nil {
		response := helper.BuildErrorResponse("Updates menu failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Updates menu successfull!", menu)
		c.JSON(http.StatusCreated, response)
	}
}

func DeleteMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var menu models.Menu
	err := config.DB.First(&menu, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Menu not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := config.DB.Delete(&menu).Error; err != nil {
		response := helper.BuildErrorResponse("Deleted menu failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Deleted menu successfull!", menu)
		c.JSON(http.StatusOK, response)
	}
}
