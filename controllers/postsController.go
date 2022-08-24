package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heyjdp/go-gin-gorm-crud/initializers"
	"github.com/heyjdp/go-gin-gorm-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data from request body
	var req struct {
		Title string
		Body  string
	}
	c.Bind(&req)

	// Create post
	post := models.Post{Title: req.Title, Body: req.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		// log.Fatal("Error adding post to DB") // log.Fatal is too hard
		c.Status(400)
		return
	} else {
		fmt.Println("Inserted post =", post.ID)
		fmt.Println("Rows affected =", result.RowsAffected)
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	} else {
		fmt.Println("Rows affected =", result.RowsAffected)
	}

	// Respond with them all
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id from URL
	var id = c.Param("id")

	// Get the posts
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	} else {
		fmt.Println("Rows affected =", result.RowsAffected)
	}

	// Respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id from the URL
	var id = c.Param("id")

	// Get the data from the request body
	var req struct {
		Title string
		Body  string
	}
	c.Bind(&req)

	// Fetch the existing post
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.Status(400)
		return
	} else {
		fmt.Println("Rows affected =", result.RowsAffected)
	}

	// Update the post
	result = initializers.DB.Model(&post).Updates(models.Post{Title: req.Title, Body: req.Body})
	if result.Error != nil {
		c.Status(400)
		return
	} else {
		fmt.Println("Rows affected =", result.RowsAffected)
	}

	// Respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the post id number
	var id = c.Param("id")

	// Request delete
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"deleted post": id,
	})
}
