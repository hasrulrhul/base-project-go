package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUserMenu(c *gin.Context) {
	var usermenu []models.UserMenu
	config.DB.Preload("Role").Preload("Menu").Find(&usermenu)
	response := helper.BuildResponse(true, "List of usermenu!", usermenu)
	c.JSON(http.StatusOK, response)
}

func CreateUserMenu(c *gin.Context) {
	var usermenu models.UserMenu
	if err := c.BindJSON(&usermenu); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	if err := config.DB.Create(&usermenu).Error; err != nil {
		response := helper.BuildErrorResponse("Created usermenu failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Created usermenu successfull!", usermenu)
		c.JSON(http.StatusCreated, response)
	}
}

func ShowUserMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var usermenu models.UserMenu
	err := config.DB.Preload("Role").Preload("Menu").First(&usermenu, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Usermenu not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Detail usermenu!", usermenu)
		c.JSON(http.StatusCreated, response)
	}
}

func UpdateUserMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var usermenu models.UserMenu
	err := config.DB.First(&usermenu, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Usermenu not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := c.BindJSON(&usermenu); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	if err := config.DB.Updates(&usermenu).Error; err != nil {
		response := helper.BuildErrorResponse("Updates usermenu failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Updates usermenu successfull!", usermenu)
		c.JSON(http.StatusCreated, response)
	}
}

func DeleteUserMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	var usermenu models.UserMenu
	err := config.DB.First(&usermenu, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Usermenu not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := config.DB.Delete(&usermenu).Error; err != nil {
		response := helper.BuildErrorResponse("Deleted usermenu failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Deleted usermenu successfull!", usermenu)
		c.JSON(http.StatusOK, response)
	}
}
