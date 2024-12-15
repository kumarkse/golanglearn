package main

import (
	"GoGinGonic/controller"
	"GoGinGonic/database"
	"GoGinGonic/initializers"
	"GoGinGonic/middleware"
	"GoGinGonic/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDb()
	migrate.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/post", controller.PostCreate)
	r.GET("/post", controller.PostIndex)
	r.GET("/post/:id", controller.PostShow)
	r.PUT("/post/:id", controller.PostUpdate)
	r.DELETE("/post/:id", controller.PostDelete)

	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/validate",middleware.RequireAuth, controller.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
