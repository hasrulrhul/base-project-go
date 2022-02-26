package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexOption(c *gin.Context) {
	var option []models.Option
	config.DB.Find(&option)
	response := helper.BuildResponse(true, "List of option!", option)
	c.JSON(http.StatusOK, response)
}

func CreateOption(c *gin.Context) {
	var option models.Option
	if err := c.BindJSON(&option); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	if err := config.DB.Create(&option).Error; err != nil {
		response := helper.BuildErrorResponse("Created option failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Created option successfull!", option)
		c.JSON(http.StatusCreated, response)
	}
}

func ShowOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var option models.Option
	err := config.DB.First(&option, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Option not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Detail option!", option)
		c.JSON(http.StatusCreated, response)
	}
}

func UpdateOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var option models.Option
	err := config.DB.First(&option, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Option not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := c.BindJSON(&option); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	if err := config.DB.Updates(&option).Error; err != nil {
		response := helper.BuildErrorResponse("Updates option failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Updates option successfull!", option)
		c.JSON(http.StatusCreated, response)
	}
}

func DeleteOption(c *gin.Context) {
	id := c.Params.ByName("id")
	var option models.Option
	err := config.DB.First(&option, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Option not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := config.DB.Delete(&option).Error; err != nil {
		response := helper.BuildErrorResponse("Deleted option failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Deleted option successfull!", option)
		c.JSON(http.StatusOK, response)
	}
}
