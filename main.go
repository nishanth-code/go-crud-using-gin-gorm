package main

import (
	"go-rest-api/controllers"
	"go-rest-api/initializers"

	"github.com/gin-gonic/gin"
)

func  init() {
	initializers.LoadVariables()
	initializers.ConnectToDb()

}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostADD)
	r.PUT("/post/:id",controllers.PostUpdate)
	r.DELETE("/post/:id",controllers.Postdelete)
	r.GET("/post",controllers.Postretrive)
	r.GET("/post/:id",controllers.PostGet)

	r.Run() 
}