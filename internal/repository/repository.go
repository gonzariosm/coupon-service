package repository

import "coupon_service/internal/entity"

type Repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}
