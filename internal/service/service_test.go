package service

import (
	"coupon_service/internal/entity"
	"coupon_service/internal/repository/memdb"
	"testing"
)

func TestCreateCoupon(t *testing.T) {
	repo := memdb.New()
	svc := New(repo)

	coupon, err := svc.CreateCoupon(10, "CODE10", 50)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if coupon.Discount != 10 {
		t.Errorf("expected discount to be 10, got %d", coupon.Discount)
	}

	if coupon.Code != "CODE10" {
		t.Errorf("expected code to be CODE10, got %s", coupon.Code)
	}

	if coupon.MinBasketValue != 50 {
		t.Errorf("expected minBasketValue to be 50, got %d", coupon.MinBasketValue)
	}
}

func TestApplyCoupon(t *testing.T) {
	repo := memdb.New()
	svc := New(repo)

	// Create test coupon
	_, err := svc.CreateCoupon(10, "CODE10", 50)
	if err != nil {
		t.Fatalf("failed to create coupon: %v", err)
	}

	// Success case
	basket := entity.Basket{Value: 100}
	updatedBasket, err := svc.ApplyCoupon(basket, "CODE10")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if updatedBasket.AppliedDiscount != 10 {
		t.Errorf("expected applied discount to be 10, got %d", updatedBasket.AppliedDiscount)
	}

	if !updatedBasket.ApplicationSuccessful {
		t.Errorf("expected application successful to be true")
	}

	// Fail case: basket value is less than minimum value
	basket = entity.Basket{Value: 40}
	_, err = svc.ApplyCoupon(basket, "CODE10")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGetCoupons(t *testing.T) {
	repo := memdb.New()
	svc := New(repo)

	// Create test coupons
	_, err := svc.CreateCoupon(10, "CODE10", 50)
	if err != nil {
		t.Fatalf("failed to create coupon: %v", err)
	}
	_, err = svc.CreateCoupon(20, "CODE20", 100)
	if err != nil {
		t.Fatalf("failed to create coupon: %v", err)
	}

	// Get coupons
	coupons, err := svc.GetCoupons([]string{"CODE10", "CODE20"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(coupons) != 2 {
		t.Errorf("expected 2 coupons, got %d", len(coupons))
	}

	if coupons[0].Code != "CODE10" || coupons[1].Code != "CODE20" {
		t.Errorf("coupons do not match the expected values")
	}
}
