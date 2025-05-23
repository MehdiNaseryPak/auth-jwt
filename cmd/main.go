package main

import (
	"github.com/MehdiNaseryPak/auth/controllers"
	"github.com/MehdiNaseryPak/auth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.POST("/register",controllers.Register)
	router.Run()
}
