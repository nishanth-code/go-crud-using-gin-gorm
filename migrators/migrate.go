package main

import (
	"go-rest-api/initializers"
	 "go-rest-api/models"
)

func init() {
	// initializers.LoadVariables()
	initializers.ConnectToDb()

}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}