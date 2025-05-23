package controllers

import (
	"github.com/MehdiNaseryPak/auth/helpers"
	"github.com/MehdiNaseryPak/auth/initializers"
	"github.com/MehdiNaseryPak/auth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var body struct {
		Username string
		Password string
		Email    string
	}

	c.Bind(&body)
	var user []models.User
	userExists := initializers.DB.Where("username = ?", body.Username).First(&user)

	// check username exists or not
	if userExists.Error == nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "این نام کاربری وجود دارد",
			"data":    userExists,
		})
	}

	// create user

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	userCreated := models.User{Username: body.Username, Password: string(hashedPassword), Email: body.Email}
	result := initializers.DB.Create(&userCreated)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user created",
		"data":    userCreated,
	})
}


func Login(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)
	var user models.User
	userExists := initializers.DB.Where("username = ?", body.Username).First(&user)

	// check username exists or not
	if userExists.Error == nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "این نام کاربری وجود دارد",
			"data":    userExists,
		})
	}

	// check password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "کلمه عبور اشتباه است",
			"data":    body.Password,
		})
	}

	// create token
	token,err := helpers.GenerateJWT(user.ID)
	response := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email": user.Email,
		"token":    token,
	}
	if err != nil {
		panic(err)
	}
	
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user login",
		"data":    response,
	})
}
