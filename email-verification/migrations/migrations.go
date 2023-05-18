package main

import (
	"email-verification/initializers"
	"email-verification/models"
	"email-verification/utils"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {
	fmt.Println("Migrating database")
	err := initializers.DB.AutoMigrate(&models.User{})
	utils.CheckError(err, "Error migrating database")
}
