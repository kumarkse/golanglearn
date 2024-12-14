package controller

import (
	"GoGinGonic/database"
	"GoGinGonic/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// get data from req body

	var body struct {
		Body string 
		Title string
	}

	c.Bind(&body)

	// create a post data
	post := models.Post{Title: body.Title, Body: body.Body}
	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// get the posts
	var posts [] models.Post

	database.DB.Find(&posts)

	c.JSON(200,gin.H{
		"posts" : posts,
	})

}

func PostShow(c *gin.Context){
	
	var id = c.Param("id")
	var post models.Post
	database.DB.Find(&post,id)

	c.JSON(200,gin.H{
		"post":post,
	})

}

func PostUpdate(c *gin.Context){
	// get the id of post to update
	var id =  c.Param("id")

	// get the body containing detail to be updated (create a struct that matches the body structure and then bind it)
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)


	// update the DB

	var post models.Post 
	database.DB.Find(&post,id)

	database.DB.Model(&post).Updates(models.Post{
		Title : body.Title,
		Body: body.Body,
	})
	
	
	
	
	// return the updated Post
	c.JSON(200,gin.H{
		"post":post,
	})
}

func PostDelete(c *gin.Context){
	id := c.Param("id")

	database.DB.Delete(&models.Post{},id)

	c.Status(200)

}