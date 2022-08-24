package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heyjdp/go-gin-gorm-crud/controllers"
	"github.com/heyjdp/go-gin-gorm-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PATCH("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}