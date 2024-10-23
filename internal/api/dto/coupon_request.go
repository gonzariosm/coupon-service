package dto

type CouponRequest struct {
	Codes []string `json:"codes"`
}

type Coupon struct {
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"min_basket_value"`
}
