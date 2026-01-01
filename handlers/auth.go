package handlers

import (
	"blog-platform/database"
	"blog-platform/model"
	"blog-platform/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := utils.HashPassword(input.Password)
	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}
	userResponse := model.DetailedUserResponse{
		ID:            user.ID,
		Email:         user.Email,
		Name:          user.Name,
		CreatedAt:     user.CreatedAt.Format(time.RFC3339),
		EmailVerified: false,
	}
	response := model.SuccessResponse{}
	response.Success.Status = http.StatusCreated
	response.Success.Message = "User registered successfully"
	response.Success.Data = userResponse
	c.JSON(http.StatusCreated, response)
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	userResponse := model.DetailedUserResponse{
		ID:            user.ID,
		Email:         user.Email,
		Name:          user.Name,
		CreatedAt:     user.CreatedAt.Format(time.RFC3339),
		EmailVerified: false,
	}
	response := model.SuccessResponse{}
	response.Success.Status = http.StatusOK
	response.Success.Message = "Login successful"
	response.Success.Data = model.AuthData{
		User:  userResponse,
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}
