package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stajprojesi/config"
	"stajprojesi/models"
	"time"
)

// GetRoomsType retrieves all room types
func GetRoomsType(c *gin.Context) {
	var roomTypes []models.RoomType
	if err := config.DB.Find(&roomTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda tipleri alınırken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roomTypes": roomTypes})
}

// GetRoomTypeByID retrieves a specific room type by its ID
func GetRoomTypeByID(c *gin.Context) {
	var roomType models.RoomType
	roomTypeID := c.Param("room_type_id") // ID parametresini alıyoruz

	// ID ile oda tipini bul
	if err := config.DB.First(&roomType, "id = ?", roomTypeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Belirtilen ID'ye sahip oda tipi bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roomType": roomType})
}

// CreateRoomType creates a new room type
func CreateRoomType(c *gin.Context) {
	var roomType models.RoomType
	if err := c.ShouldBindJSON(&roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}
	// Oluşturma ve güncelleme tarihlerini ayarlama
	roomType.CreateDate = time.Now()
	roomType.UpdateDate = time.Now()

	if err := config.DB.Create(&roomType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda tipi oluşturulurken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Oda tipi başarıyla oluşturuldu", "roomType": roomType})
}

// UpdateRoomType updates an existing room type
func UpdateRoomType(c *gin.Context) {
	var roomType models.RoomType
	roomTypeID := c.Param("room_type_id") // URL'den oda tipi ID'sini alıyoruz

	// ID ile oda tipini bul
	if err := config.DB.First(&roomType, "id = ?", roomTypeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Güncellenecek oda tipi bulunamadı"})
		return
	}

	// Gelen JSON verisini oda tipine bağlama
	if err := c.ShouldBindJSON(&roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	// Güncelleme tarihini ayarla
	roomType.UpdateDate = time.Now()

	// Güncelleme işlemi
	if err := config.DB.Save(&roomType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda tipi güncellenirken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Oda tipi başarıyla güncellendi", "roomType": roomType})
}

// DeleteRoomType deletes a room type
func DeleteRoomType(c *gin.Context) {
	roomTypeID := c.Param("room_type_id") // URL'den oda tipi ID'sini alıyoruz

	// Oda tipini ID ile silme işlemi
	if err := config.DB.Delete(&models.RoomType{}, "id = ?", roomTypeID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oda tipi silinirken bir hata oluştu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Oda tipi başarıyla silindi"})
}
