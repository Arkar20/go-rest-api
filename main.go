package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// db.InitDB()
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    "restapi working",
		})
	})

	r.Run(":3000")
}
