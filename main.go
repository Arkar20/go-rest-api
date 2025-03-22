package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	ID             uint   `gorm:"primaryKey;column:id"`
	Name           string `gorm:"size:255;column:name"`
	PhNum1         string `gorm:"size:20;column:phnum1"`
	PhNum2         string `gorm:"size:20;column:phnum2"`
	Company        string `gorm:"size:255;column:company"`
	Address        string `gorm:"size:500;column:address"`
	RetailCustomer bool   `gorm:"default:false;column:retail_customer"`
	MarketplaceID  uint   `gorm:"column:marketplace_id"`
}

var db *gorm.DB

func initDB() {
	err := godotenv.Load()
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")

	fmt.Println(dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// db.AutoMigrate(&User{})
	log.Println("Connected to the database")
}

func getUsers(c *gin.Context) {
	var users []Customer
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func main() {
	initDB()
	r := gin.Default()
	r.GET("/users", getUsers)

	r.Run(":8080")
}
