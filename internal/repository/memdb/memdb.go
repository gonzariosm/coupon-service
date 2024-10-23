package memdb

import (
	"coupon_service/internal/entity"
	"coupon_service/internal/repository"
	"fmt"
)

type Repository struct {
	entries map[string]entity.Coupon
}

func New() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon), // init entries
	}
}

// FindByCode looks up a coupon by its code in the in-memory store.
// Returns the coupon if found or an error if the coupon does not exist.
func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("coupon not found")
	}
	return &coupon, nil
}

// Save stores a new coupon in the in-memory repository.
// Returns an error if the coupon could not be saved.
func (r *Repository) Save(coupon entity.Coupon) error {
	r.entries[coupon.Code] = coupon
	return nil
}

// Implement the repository interface
var _ repository.Repository = &Repository{}
