package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{DB: db}
}

// GetCustomers handles the route to get all customers
func (controller *Controller) GetCustomers(ctx *gin.Context) {
	var customers []Customer

	if err := controller.DB.Find(&customers).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, ToCustomerDTOArray(customers))
}
