package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"stajprojesi/config"
	"stajprojesi/models"
)

// GetCustomer retrieves a specific customer by ID
func GetCustomer(context *gin.Context) {
	var customer models.Customer
	customerID := context.Param("customer_id")

	if err := config.DB.First(&customer, "customer_id = ?", customerID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"customers": customer})
}

// GetCustomers lists all customers from the database
func GetCustomers(context *gin.Context) {
	var customers []models.Customer
	// Fetch all customers from the database
	if err := config.DB.Find(&customers).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}
	// Return the list of customers
	context.JSON(http.StatusOK, gin.H{"customers": customers})
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the given password matches the hashed password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CustomerSignUp creates a new customer in the customers table
func CustomerSignUp(context *gin.Context) {
	var customer models.Customer
	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Hash the customer's password before saving to the database
	hashedPassword, err := HashPassword(customer.PasswordHash)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	customer.PasswordHash = hashedPassword

	if err := config.DB.Create(&customer).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Customer registered successfully"})
}

// CustomerLogin authenticates a customer by email and password
func CustomerLogin(context *gin.Context) {
	var request models.Customer
	var customer models.Customer
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Find the customer by email
	if err := config.DB.Where("email = ?", request.Email).First(&customer).Error; err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Verify the password
	if err := VerifyPassword(customer.PasswordHash, request.PasswordHash); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Customer login successful", "customers": customer})
}
func UpdateCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	var customer models.Customer

	// Check if the customer exists
	if err := config.DB.First(&customer, customerID).Error; err != nil { // Düzeltme: models.DB yerine config.DB kullanılıyor
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Bind the JSON payload to the customer model
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	// Save the updated customer
	if err := config.DB.Save(&customer).Error; err != nil { // Düzeltme: models.DB yerine config.DB kullanılıyor
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully", "data": customer})
}

// DeleteCustomer handles deleting a customer
func DeleteCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	var customer models.Customer

	// Check if the customer exists
	if err := config.DB.First(&customer, customerID).Error; err != nil { // Düzeltme: models.DB yerine config.DB kullanılıyor
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Delete the customer
	if err := config.DB.Delete(&customer).Error; err != nil { // Düzeltme: models.DB yerine config.DB kullanılıyor
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
