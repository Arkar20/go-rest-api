package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, continuing with environment variables")
	}

	// Get database URL from environment
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not set in environment")
	}

	// Connect to PostgreSQL database using GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Optional: Run auto migration
	// DB.AutoMigrate(&User{})

	log.Println("Connected to the database")
}
