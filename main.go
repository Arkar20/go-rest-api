package main

import (
	"restapi/customer"
	"restapi/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(cors.Default())

	customer.RegisterRoutes(r, customer.NewController())

	r.Run(":3000")
}
