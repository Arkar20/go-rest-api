package customer

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes will register all the customer-related routes
func RegisterRoutes(r *gin.Engine, customerController *Controller) {
	r.GET("/customers", customerController.GetCustomers)
}
