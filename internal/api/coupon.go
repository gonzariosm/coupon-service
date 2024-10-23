package api

import (
	"coupon_service/internal/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Apply handles the POST request to apply a coupon to a basket.
// It binds the request body to an ApplicationRequest and applies the coupon using the service.
// Returns the updated basket or an error if the coupon could not be applied.
func (a *API) Apply(c *gin.Context) {
	var apiReq dto.ApplicationRequest
	// Error JSON binding
	if !bindJSON(c, &apiReq) {
		return
	}

	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basket)
}

// Create handles the POST request to create a new coupon.
// It validates the input and creates a coupon using the service.
// Returns the created coupon or an error if the operation fails.
func (a *API) Create(c *gin.Context) {
	var apiReq dto.Coupon
	// Error JSON binding
	if !bindJSON(c, &apiReq) {
		return
	}

	// Validate fields
	if apiReq.Code == "" || apiReq.Discount <= 0 || apiReq.MinBasketValue <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coupon parameters {code, discount, min_basket_value}"})
		return
	}

	coupon, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return status created and coupon data
	c.JSON(http.StatusCreated, coupon)
}

// Get handles the GET request to retrieve coupons by their codes.
// It returns the coupons found or an error if no valid coupons are found.
func (a *API) Get(c *gin.Context) {
	var apiReq dto.CouponRequest
	// Error JSON binding
	if !bindJSON(c, &apiReq) {
		return
	}

	// Validate if user sent codes
	if len(apiReq.Codes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No coupon codes provided"})
		return
	}

	// Get coupons
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil && len(coupons) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		// 207 Multi-Status and details
		c.JSON(http.StatusMultiStatus, gin.H{
			"message": "Some coupons were not found",
			"error":   err.Error(),
			"coupons": coupons,
		})
		return
	}

	// Return coupons
	c.JSON(http.StatusOK, coupons)
}

// Helper function to bind JSON and handle errors generically
func bindJSON[T any](c *gin.Context, obj *T) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return false
	}
	return true
}
