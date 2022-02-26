package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/helper"
	"base-project-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {
	var user []models.User
	config.DB.Preload("Role").Find(&user)
	response := helper.BuildResponse(true, "List of user!", user)
	c.JSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed enkripsi")
		return
	}
	user.Password = hashedPassword
	if err := config.DB.Create(&user).Error; err != nil {
		response := helper.BuildErrorResponse("Created user failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Created user successfull!", user)
		c.JSON(http.StatusCreated, response)
	}
}

func ShowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.Preload("Role").First(&user, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("User not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Detail user!", user)
		c.JSON(http.StatusCreated, response)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.Preload("Role").First(&user, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("User not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := c.BindJSON(&user); err != nil {
		for _, v := range c.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"error": v.Error()})
			return
		}
	}
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed enkripsi")
		return
	}
	user.Password = hashedPassword
	if err := config.DB.Updates(&user).Error; err != nil {
		response := helper.BuildErrorResponse("Updates user failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Updates user successfull!", user)
		c.JSON(http.StatusCreated, response)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		response := helper.BuildErrorResponse("User not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		response := helper.BuildErrorResponse("Deleted user failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Deleted user successfull!", user)
		c.JSON(http.StatusOK, response)
	}
}

func UploadUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
