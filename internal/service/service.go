package service

import (
	"coupon_service/internal/entity"
	"coupon_service/internal/repository"

	"fmt"

	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

// ApplyCoupon applies a discount to a basket if the coupon code is valid.
// It checks if the basket's value meets the coupon's minimum basket value requirement.
// Returns the updated basket with the applied discount or an error if conditions are not met.
func (s Service) ApplyCoupon(basket entity.Basket, code string) (*entity.Basket, error) {
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if basket.Value <= 0 {
		return nil, fmt.Errorf("basket value is not valid, no discount applied")
	}

	if basket.Value < coupon.MinBasketValue {
		return nil, fmt.Errorf("basket value is below the minimum required: %d", coupon.MinBasketValue)
	}

	basket.AppliedDiscount = coupon.Discount
	basket.ApplicationSuccessful = true
	return &basket, nil
}

// CreateCoupon creates a new coupon with a discount, code, and minimum basket value.
// The coupon is stored in the repository and returned on success.
// Returns the created coupon or an error if the operation fails.
func (s Service) CreateCoupon(discount int, code string, minBasketValue int) (*entity.Coupon, error) {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return nil, err
	}
	return &coupon, nil
}

// GetCoupons retrieves a list of coupons by their codes.
// It returns the coupons that were found and an error if any of the provided codes were not found.
func (s Service) GetCoupons(codes []string) ([]entity.Coupon, error) {
	coupons := make([]entity.Coupon, 0, len(codes))
	var errorsList []error

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			errorsList = append(errorsList, fmt.Errorf("coupon not found for code: %s, index: %d", code, idx))
			continue // fix to do not add no existing coupons
		}
		coupons = append(coupons, *coupon)
	}

	if len(errorsList) > 0 {
		return coupons, errors.Join(errorsList...)
	}

	return coupons, nil
}
