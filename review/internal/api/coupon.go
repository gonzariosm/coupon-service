package api

import (
	. "coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply(c *gin.Context) {
	var apiReq ApplicationRequest
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

func (a *API) Create(c *gin.Context) {
	var apiReq Coupon
	// Error JSON binding
	if !bindJSON(c, &apiReq) {
		return
	}

	// Validate fields (basic validation)
	if apiReq.Code == "" || apiReq.Discount <= 0 || apiReq.MinBasketValue <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coupon parameters {code, discount, minBasketValue}"})
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// TODO: modify the service to return the new object with the uuid
	c.Status(http.StatusOK)
}

func (a *API) Get(c *gin.Context) {
	var apiReq CouponRequest
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
