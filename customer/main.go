package customer

import (
	"restapi/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes will register all the customer-related routes
func RegisterRoutes(r *gin.Engine, customerController *Controller) {
	r.GET("/customers", customerController.Index)
	r.GET("/customers/:id", customerController.Show)
	r.POST("/customers/sign-in", customerController.SignIn)

	protected := r.Group("/customers")

	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("customer_id")
			c.JSON(200, gin.H{"message": "Welcome!", "userID": userID})
		})
	}
}
