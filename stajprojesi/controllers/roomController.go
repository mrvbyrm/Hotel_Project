package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stajprojesi/config"
	"stajprojesi/models"
)

// GetRoomByID gets a specific room by ID
func GetRoomByID(c *gin.Context) {
	roomID := c.Param("id")
	var room models.Room
	if err := config.DB.First(&room, "id = ?", roomID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Belirtilen ID'ye sahip oda bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, room)
}

// CreateRoom creates a new room
func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz oda verisi"})
		return
	}
	if err := config.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda oluşturulurken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusCreated, room)
}

// UpdateRoom updates an existing room
func UpdateRoom(c *gin.Context) {
	roomID := c.Param("id")
	var room models.Room
	if err := config.DB.First(&room, "id = ?", roomID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Güncellenecek oda bulunamadı"})
		return
	}
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz güncelleme verisi"})
		return
	}
	if err := config.DB.Save(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda güncellenirken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusOK, room)
}

// DeleteRoom deletes a room
func DeleteRoom(c *gin.Context) {
	roomID := c.Param("id")
	if err := config.DB.Delete(&models.Room{}, "id = ?", roomID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda silinirken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Oda başarıyla silindi"})
}

// GetRooms retrieves rooms with optional filtering
func GetRooms(c *gin.Context) {
	var rooms []models.Room
	roomType := c.Query("type_name")
	query := config.DB.Model(&models.Room{})
	if roomType != "" {
		query = query.Where("type_name = ?", roomType)
	}
	if err := query.Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Odalar alınırken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusOK, rooms)
}
