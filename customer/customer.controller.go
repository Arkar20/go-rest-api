package customer

import (
	"net/http"
	"os"
	"restapi/db"
	"restapi/helper"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	CustomerId uint `json:"customer_id"`
	jwt.StandardClaims
}

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
		ctx.JSON(http.StatusBadRequest, "Invalid Param ID")
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

func (controller *Controller) SignIn(ctx *gin.Context) {

	var customer Customer

	if err := db.DB.Where("phnum1 = ?", ctx.PostForm("phnum")).First(&customer).Error; err != nil {
		ctx.JSON(http.StatusNotFound, "Customer Not Found")
		return
	}

	// Generate JWT token
	access_token, err := generateJWTToken(customer.ID, 15*time.Minute)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error generating Access token")
		return
	}

	refresh_token, err := generateJWTToken(customer.ID, 7*24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
		return
	}

	data := gin.H{
		"access_token":  access_token,
		"refresh_token": refresh_token,
		"user":          ToCustomerDTO(customer),
	}

	ctx.JSON(http.StatusOK, helper.Response{
		Success: true,
		Message: "Sign In Successful",
		Data:    data,
	})
}

// Function to generate JWT token
func generateJWTToken(customerID uint, duration time.Duration) (string, error) {
	// Define the token expiration time
	expirationTime := time.Now().Add(duration) // Token expires in 24 hours

	// Create the JWT claims, which includes the customer ID and expiration time
	claims := &Claims{
		CustomerId: customerID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   string(customerID),
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Secret key to sign the token (keep this secret safe)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
