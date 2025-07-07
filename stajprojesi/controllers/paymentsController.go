package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stajprojesi/config"
	"stajprojesi/models"
	"strconv"
)

// GetPayment handles GET requests to fetch all payments
func GetPayment(c *gin.Context) {
	// Fetch all payments from the database
	payments, err := models.GetPayment(config.DB) // Correct function name
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch payments"})
		return
	}
	c.JSON(http.StatusOK, payments)
}

// GetPaymentByID handles GET requests to fetch a single payment by ID
func GetPaymentByID(c *gin.Context) {
	paymentIDStr := c.Param("id")
	paymentID, err := strconv.ParseUint(paymentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}
	// Veritabanından ödeme bilgilerini al
	payment, err := models.GetPaymentByID(config.DB, uint(paymentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

// CreatePayment handles POST requests to create a new payment
func CreatePayment(c *gin.Context) {
	var newPayment models.Payment
	if err := c.ShouldBindJSON(&newPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Add new payment record to the database
	if err := models.CreatePayment(config.DB, newPayment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create payment"})
		return
	}
	c.JSON(http.StatusCreated, newPayment)
}

// RefundPayment handles POST requests to refund a payment
func RefundPayment(c *gin.Context) {
	var refundRequest struct {
		PaymentID string  `json:"id"`
		Amount    float64 `json:"amount"` // Bu örnek için amount kullanılıyor, ancak refund işlemi için kullanılmıyor
	}
	if err := c.ShouldBindJSON(&refundRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	paymentIDStr := refundRequest.PaymentID
	paymentID, err := strconv.ParseUint(paymentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}
	if err := models.RefundPayment(config.DB, uint(paymentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to process refund"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Refund processed successfully"})
}
