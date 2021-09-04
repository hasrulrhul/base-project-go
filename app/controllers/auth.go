package controllers

import (
	"base-project-go/app/models"
	"base-project-go/config"
	"base-project-go/service"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	name     = "name"
	username = "username"
	email    = "email"
	role     = "role"
)

func Login(c *gin.Context) {
	var u models.LoginCredentials
	if err := c.BindJSON(&u); err != nil {
		panic(err)
	}
	var user models.User
	err := config.DB.Preload("Role").Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
	} else {
		// cek password
		match := service.CheckPasswordHash(u.Password, user.Password)
		if match {
			session := sessions.Default(c)
			// In real world usage you'd set this to the users
			session.Set(name, user.Name)
			session.Set(username, user.Username)
			session.Set(email, user.Email)
			session.Set(role, user.Role.Role)
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}
			// tokens, err := service.GenerateTokenPair()
			// if err != nil {
			// 	c.JSON(http.StatusBadRequest, "failed generate token")
			// 	return
			// }
			// c.JSON(http.StatusBadRequest, gin.H{"data": user, "token": tokens, "message": "login success"})
			c.JSON(http.StatusBadRequest, gin.H{"data": user, "message": "login success"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "password wrong"})
		}
	}

	// result, err := config.DB.Where("email = ?", user.Email).First(&user)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "user not found")
	// }

	// c.JSON(http.StatusOK, tokens)
	// c.JSON(http.StatusOK, service.Response(result, c, "", 0))

	// c.JSON(http.StatusOK, gin.H{"user": user.Username, "pass": user.Password})
	// c.JSON(http.StatusOK, gin.H{"user": user.Username, "pass": user.Password})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed enkripsi")
		return
	}
	user.Password = hashedPassword
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "register success")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(name)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(name)
	session.Delete(username)
	session.Delete(email)
	session.Delete(role)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func GetSession(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get(name)
	username := session.Get(username)
	email := session.Get(email)
	role := session.Get(role)
	c.JSON(http.StatusOK, gin.H{"name": name, "username": username, "email": email, "role": role})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
