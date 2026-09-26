package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"stajprojesi/models"
)

var DB *gorm.DB // Publicly accessible database connection

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Connect establishes a database connection, performs migrations, and checks the connection status
func Connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Istanbul",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "superuser"),
		getEnv("DB_NAME", "postgres"),
		getEnv("DB_PORT", "5432"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlDB: %v", err)
	}

	// Ping to check if the database connection is alive
	start := time.Now()
	for {
		if err := sqlDB.Ping(); err == nil {
			break
		}
		if time.Since(start) > 10*time.Second {
			return fmt.Errorf("database connection timeout")
		}
		time.Sleep(1 * time.Second) // Delay between retries
	}
	log.Println("Database connection established")

	// Automatic migration
	if err := DB.AutoMigrate(
		&models.Admin{},
		&models.Room{},
		&models.RoomType{},
		&models.Reservation{},
		&models.Customer{},
		&models.Payment{},
		&models.User{},
	); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}
	return nil
}

func Setup() {
	// Setup function - add configuration or initialization code if needed
	fmt.Println("Config setup")
}
