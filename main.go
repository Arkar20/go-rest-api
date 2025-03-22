package main

import (
	"net/http"
	"restapi/db"

	"github.com/gin-gonic/gin"
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

func getUsers(c *gin.Context) {
	var users []Customer
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func main() {
	db.InitDB()
	r := gin.Default()
	r.GET("/users", getUsers)

	r.Run(":8080")
}
