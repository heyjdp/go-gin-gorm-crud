package main

import (
	"github.com/heyjdp/go-gin-gorm-crud/initializers"
	"github.com/heyjdp/go-gin-gorm-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
