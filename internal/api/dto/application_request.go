package dto

import "coupon_service/internal/entity"

type ApplicationRequest struct {
	Code   string        `json:"code"`
	Basket entity.Basket `json:"basket"`
}
