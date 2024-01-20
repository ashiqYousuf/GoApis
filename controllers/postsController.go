package controllers

import (
	"github.com/ashiqYousuf/GoApisStructure/initializers"
	"github.com/ashiqYousuf/GoApisStructure/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	post := models.Post{}
	err := c.Bind(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  []string{"something went wrong"},
		})
		return
	}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(201, gin.H{
		"success": true,
		"data": gin.H{
			"post": post,
		},
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"results": len(posts),
		"data": gin.H{
			"posts": posts,
		},
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	tx := initializers.DB.First(&post, id)

	if tx.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"errors":  []string{"record not found"},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"post": post,
		},
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	tx := initializers.DB.First(&post, id)

	if tx.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"errors":  []string{"record not found"},
		})
		return
	}
	err := c.Bind(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  []string{"something went wrong"},
		})
		return
	}
	tx = initializers.DB.Model(&post).Updates(models.Post{
		Title: post.Title,
		Body:  post.Body,
	})

	if tx.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  []string{"post not updated"},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"post": post,
		},
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	tx := initializers.DB.First(&post, id)

	if tx.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"errors":  []string{"record not found"},
		})
		return
	}

	tx = initializers.DB.Delete(&models.Post{}, id)
	if tx.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  []string{"record not deleted"},
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
	})
}
