package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stajprojesi/config"
	"stajprojesi/models"
)

func GetReservations(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		c.Abort()
		return
	}

	reservations, err := models.GetReservationsByUserID(config.DB, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch reservations"})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func CreateReservation(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		c.Abort()
		return
	}
	var newReservation models.Reservation
	if err := c.ShouldBindJSON(&newReservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Doğru struct alanını kullanın
	newReservation.UserID = userID.(int)

	if err := models.CreateReservation(config.DB, newReservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create reservation"})
		return
	}
	c.JSON(http.StatusCreated, newReservation)
}

func GetReservationByID(c *gin.Context) {
	reservationID := c.Param("id")
	reservation, err := models.GetReservationByID(config.DB, reservationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservation)
}
func UpdateReservation(c *gin.Context) {
	reservationID := c.Param("id")
	var updatedReservation models.Reservation
	if err := c.ShouldBindJSON(&updatedReservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.UpdateReservation(config.DB, reservationID, updatedReservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update reservation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Reservation updated successfully"})
}

func DeleteReservation(c *gin.Context) {
	reservationID := c.Param("id")
	if err := models.DeleteReservation(config.DB, reservationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete reservation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Reservation deleted successfully"})
}
