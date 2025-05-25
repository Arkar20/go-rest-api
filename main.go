package main

import (
	"restapi/customer"
	"restapi/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.New()

	r.Use(gzip.Gzip(gzip.BestSpeed))
	
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://frontend.local.com"}

	r.Use(cors.New(config))

	customer.RegisterRoutes(r, customer.NewController())

	r.Run(":3000")
}
