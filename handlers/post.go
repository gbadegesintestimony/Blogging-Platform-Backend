package handlers

import (
	"blog-platform/database"
	"blog-platform/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

func CreatePost(c *gin.Context) {
	userID := c.GetUint("userID")
	var input PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post := model.Post{
		Title:      input.Title,
		Content:    input.Content,
		UserID:     userID,
		CategoryID: input.CategoryID,
	}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create post"})
		return
	}
	c.JSON(201, gin.H{"message": "Post created successfully", "post": post})
}

func GetPosts(c *gin.Context) {
	var posts []model.Post
	if err := database.DB.Preload("comments").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not retrieve posts"})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var post model.Post
	if err := database.DB.Preload("comments").First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	var input PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}
	if post.UserID != userID {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}
	post.Title = input.Title
	post.Content = input.Content
	post.CategoryID = input.CategoryID
	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not update post"})
		return
	}
	c.JSON(200, gin.H{"message": "Post updated successfully", "post": post})
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}
	if post.UserID != userID {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}
	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not delete post"})
		return
	}
	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
