package itemPromo

import (
	"context"
	itemPromoEntity "go-itempromo/internal/entity/itemPromo"

	"github.com/pkg/errors"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetItemPromo(ctx context.Context) ([]itemPromoEntity.ItemPromo, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data Data
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data) Service {
	// Assign variable dari parameter ke object
	return Service{
		data: data,
	}
}

// GetItemPromo ..
func (s Service) GetItemPromo(ctx context.Context) ([]itemPromoEntity.ItemPromo, error) {
	// Panggil method GetAllUsers di data layer user
	skeleton, err := s.data.GetItemPromo(ctx)
	// Error handling
	if err != nil {
		return skeleton, errors.Wrap(err, "[SERVICE][GetItemPromo]")
	}
	// Return users array
	return skeleton, err
}
