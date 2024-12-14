package main

import (
	"GoGinGonic/controller"
	"GoGinGonic/database"
	"GoGinGonic/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDb()

}
func main() {
	r := gin.Default()
	r.POST("/post", controller.PostCreate)
	
	r.GET("/post", controller.PostIndex)
	
	r.GET("/post/:id", controller.PostShow)

	r.PUT("/post/:id", controller.PostUpdate)

	r.DELETE("/post/:id", controller.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
