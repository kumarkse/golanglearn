package migrate

import (
	"GoGinGonic/database"
	"GoGinGonic/initializers"
	"GoGinGonic/models"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDb()
}

// func main() {
// 	database.DB.AutoMigrate(&models.Post{})
// }

func SyncDatabase() {
	database.DB.AutoMigrate(&models.User{})
}
