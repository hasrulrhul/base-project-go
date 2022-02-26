package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRole(c *gin.Context) {
	var role []models.Role
	config.DB.Find(&role)
	response := helper.BuildResponse(true, "List of role!", role)
	c.JSON(http.StatusOK, response)
}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.BindJSON(&role); err != nil {
		panic(err)
	}
	if err := config.DB.Create(&role).Error; err != nil {
		response := helper.BuildErrorResponse("Created role failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Created role successfull!", role)
		c.JSON(http.StatusCreated, response)
	}
}

func ShowRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Role not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Detail role!", role)
		c.JSON(http.StatusCreated, response)
	}
}

func UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Role not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := c.BindJSON(&role); err != nil {
		panic(err)
	}
	if err := config.DB.Updates(&role).Error; err != nil {
		response := helper.BuildErrorResponse("Updates role failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Updates role successfull!", role)
		c.JSON(http.StatusCreated, response)
	}
}

func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("Role not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := config.DB.Delete(&role).Error; err != nil {
		response := helper.BuildErrorResponse("Deleted role failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Deleted role successfull!", role)
		c.JSON(http.StatusOK, response)
	}
}
