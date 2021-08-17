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
	if err := c.BindJSON(&role); err != nil {
		panic(err)
	}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func ShowRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
	} else {
		c.JSON(http.StatusOK, role)
	}
}

func UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := c.BindJSON(&role); err != nil {
		panic(err)
	}
	if err := config.DB.Updates(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	err := config.DB.First(&role, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
