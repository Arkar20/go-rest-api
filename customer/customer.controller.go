package customer

import (
	"net/http"
	"restapi/db"
	"restapi/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// GetCustomers handles the route to get all customers
func (controller *Controller) Index(ctx *gin.Context) {
	var customers []Customer

	if err := db.DB.Scopes(helper.Paginate(ctx.Request)).Find(&customers).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, ToCustomerDTOArray(customers))
}

// GetCustomer handles the route to get single customer
func (controller *Controller) Show(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if (err) != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	var customer Customer = Customer{
		ID: uint(id),
	}

	if err := db.DB.First(&customer).Error; err != nil {
		ctx.JSON(http.StatusNotFound, "Customer Not Found")
		return
	}

	ctx.JSON(http.StatusOK, ToCustomerDTO(customer))
}
