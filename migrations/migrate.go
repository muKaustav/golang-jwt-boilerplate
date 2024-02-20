package main

import (
	"fmt"
	"golang/initializers"
	"golang/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Author{})

	if err != nil {
		fmt.Println("Failed to migrate database:", err)
	} else {
		fmt.Println("Database migrated")
	}
}
