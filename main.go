package main

import (
	"restapi/customer"
	"restapi/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	customer.RegisterRoutes(r, customer.NewController(db.DB))

	r.Run(":8080")
}
