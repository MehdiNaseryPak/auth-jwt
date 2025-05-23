package main

import (
	"github.com/MehdiNaseryPak/auth/initializers"
	"github.com/MehdiNaseryPak/auth/models"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectDB()
}
func main() {

	initializers.DB.AutoMigrate(&models.User{})
}
