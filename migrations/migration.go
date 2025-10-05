package main

import (
	"gin-introduction/infra"
	"gin-introduction/models"
)

func main() { 
	infra.Initializer()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}