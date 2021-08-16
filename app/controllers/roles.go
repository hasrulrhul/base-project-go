package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRole(c *gin.Context) {
	var role []models.Role
	config.DB.Find(&role)
	c.JSON(http.StatusOK, role)
}

func CreateRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	config.DB.Create(&role)
	c.JSON(http.StatusOK, role)
}

func ShowRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	config.DB.First(&role, id)
	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Record not found")
	}
	c.BindJSON(&role)
	config.DB.Updates(&role)
	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Record not found")
	}
	config.DB.Delete(&role)
}
