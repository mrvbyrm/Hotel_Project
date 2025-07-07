package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stajprojesi/config"
	"stajprojesi/helpers"
	"stajprojesi/helpers/utils"
	"stajprojesi/models"
	"strconv"
	"time"
)

func GetUsers(context *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	// Return the users
	context.JSON(http.StatusOK, gin.H{"users": users})
}
func GetUser(context *gin.Context) {
	var user models.User
	userID := context.Param("user_id") // Assuming the ID is passed as a URL parameter
	// Find the user by ID
	if err := config.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// Return the user
	context.JSON(http.StatusOK, gin.H{"user": user})
}
func SignUp(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Email kontrolü
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}
	// Şifreyi hash'leme
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.PasswordHash = hashedPassword
	// Token oluşturma
	tokens, refreshToken, err := helpers.GenerateAllTokens(user.Email, user.FirstName, user.LastName, strconv.Itoa(user.ID), user.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}
	user.Token = tokens
	user.RefreshToken = refreshToken

	// Yeni kullanıcıyı kaydetme
	if err := config.DB.Create(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Başarılı kayıt sonucu
	context.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "tokens": tokens})
}

func Login(context *gin.Context) {
	var request models.User
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Find the user by email
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Verify the password
	if err := utils.VerifyPassword(user.PasswordHash, request.Password); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Generate new tokens
	tokens, refreshToken, err := helpers.GenerateAllTokens(user.Email, user.FirstName, user.LastName, strconv.Itoa(user.ID), user.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}
	// Update tokens in the database
	user.Token = tokens
	user.RefreshToken = refreshToken
	user.UpdatedAt = time.Now()

	if err := config.DB.Save(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tokens"})
		return
	}
	// Return success with user and tokens
	context.JSON(http.StatusOK, gin.H{
		"message": "User login successful",
		"user":    user,
		"tokens":  tokens,
	})
}
func UpdateUser(c *gin.Context) {
	userID := c.Param("user_id")
	var user models.User

	// Check if the user exists
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind the JSON payload to the user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	// Save the updated user
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": user})
}

// DeleteUser handles deleting a user
func DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")
	var user models.User

	// Check if the user exists
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Logic for handling logout, such as invalidating the token
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
