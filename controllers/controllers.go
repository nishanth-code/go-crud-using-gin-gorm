package controllers

import (
	"go-rest-api/initializers"
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func PostADD(c *gin.Context) {
	var Requestbody struct{
		Title string
		Body string
	}
	c.Bind(&Requestbody)

	post := model.Post{Title: Requestbody.Title,Body: Requestbody.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post":post,
	})
}

func Postretrive(c *gin.Context){
	var posts []model.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts":posts,
	})
}

func PostGet(c *gin.Context){
	id := c.Param("id")
	var post model.Post
	initializers.DB.First(&post,id)

	c.JSON(200, gin.H{
		"post":post,
	})

}

func PostUpdate(c *gin.Context){
	id := c.Param("id")
	var Requestbody struct{
		Title string
		Body string
	}
	c.Bind(&Requestbody)

	var post model.Post
	initializers.DB.First(&post,id) 

	initializers.DB.Model(&post).Updates( model.Post{Title: Requestbody.Title,Body: Requestbody.Body})

	c.JSON(200, gin.H{
		"post":post,
	})
	
}

func Postdelete(c *gin.Context){
	id := c.Param("id")

	initializers.DB.Delete(&model.Post{},id)

	c.JSON(200, gin.H{
		"msg":"deleted sucessfully",
	})
}