package controllers

import (
	"base-project-go/app/models"
	"base-project-go/app/request"
	"base-project-go/config"
	"base-project-go/helper"
	"base-project-go/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginResponse token response
type LoginResponse struct {
	ID       uint
	Name     string
	Username string
	Email    string
	Role     string
	RoleName string
	// Token        string `json:"token"`
	AccessToken string `json:"access_token"`
	// RefreshToken string `json:"refresh_token"`
}

const (
	name     = "name"
	username = "username"
	email    = "email"
	role     = "role"
)

func Login(c *gin.Context) {
	var u request.LoginValidation
	if err := c.BindJSON(&u); err != nil {
		panic(err)
	}
	var user models.User
	err := config.DB.Preload("Role").Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		response := helper.BuildErrorResponse("User not found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// cek password
	match := service.CheckPasswordHash(u.Password, user.Password)
	if match == false {
		response := helper.BuildErrorResponse("Password wrong", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// create session login users
	session := sessions.Default(c)
	session.Set(name, user.Name)
	session.Set(username, user.Username)
	session.Set(email, user.Email)
	session.Set(role, user.Role.Role)
	if err := session.Save(); err != nil {
		response := helper.BuildErrorResponse("Password wrong", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// generate token login
	jwtWrapper := service.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email, uint(user.ID))
	if err != nil {
		response := helper.BuildErrorResponse("Error signing token", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tokenResponse := LoginResponse{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role.Role,
		RoleName:    user.Role.Name,
		AccessToken: signedToken,
	}
	response := helper.BuildResponse(true, "login successfull!", tokenResponse)
	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to enkripsi", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user.Password = hashedPassword
	if err := config.DB.Create(&user).Error; err != nil {
		response := helper.BuildErrorResponse("Register failed", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Register successfull!", helper.EmptyObj{})
		c.JSON(http.StatusCreated, response)
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(name)
	if user == nil {
		response := helper.BuildErrorResponse("Invalid session token", "", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	session.Delete(name)
	session.Delete(username)
	session.Delete(email)
	session.Delete(role)
	if err := session.Save(); err != nil {
		response := helper.BuildErrorResponse("Failed to save session", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, "Successfully logged out!", helper.EmptyObj{})
	c.JSON(http.StatusOK, response)
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
	response := helper.BuildResponse(true, "You are logged in!", helper.EmptyObj{})
	c.JSON(http.StatusCreated, response)
}

func GenerateJWT(email, role string, c *gin.Context) (string, error) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		response := helper.BuildErrorResponse("Something Went Wrong", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	return tokenString, nil
}
