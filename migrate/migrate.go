package main

import (
	"GinModule/initializers"
	"GinModule/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	fmt.Println("Starting the application...")
	initializers.DB.AutoMigrate(&models.Task{})
}