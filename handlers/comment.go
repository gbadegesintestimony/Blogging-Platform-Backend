package handlers

import (
	"blog-platform/database"
	"blog-platform/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	Content string `json:"content" binding:"required"`
}

func AddComment(c *gin.Context) {
	userID := c.GetUint("userID")
	postId, _ := strconv.Atoi(c.Param("postid"))

	var input CommentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment := model.Comment{
		Content: input.Content,
		PostID:  uint(postId),
		UserID:  userID,
	}
	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not add comment"})
		return
	}
	c.JSON(201, gin.H{"message": "Comment added successfully", "comment": comment})
}
