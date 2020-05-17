package product

import (
	"context"
	"fmt"
	pEntity "product/internal/entity/product"
	"product/pkg/errors"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	TampilDetailReceiveByNoReceive(ctx context.Context, NoTranrc string) (pEntity.MstProduct, error)
}

type mpData interface {
	GetAllJSONMP(ctx context.Context, kode string) (pEntity.MstProduct, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	productData Data
	mpData      mpData
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(productData Data, mpData mpData) Service {
	// Assign variable dari parameter ke object
	return Service{
		productData: productData,
		mpData:      mpData,
	}
}

// TampilDetailMP ...
func (s Service) TampilDetailMP(ctx context.Context, kode string) (pEntity.MstProduct, error) {

	var (
		product pEntity.MstProduct
		err     error
	)

	product, err = s.mpData.GetAllJSONMP(ctx, kode)

	fmt.Println("test")

	// Error handling
	if err != nil {
		return product, errors.Wrap(err, "[SERVICE][TampilDetailMP]")
	}
	// Return users array
	return product, err
}

// TampilDetailReceiveByNoReceive ..
func (s Service) TampilDetailReceiveByNoReceive(ctx context.Context, NoTranrc string) (pEntity.MstProduct, error) {
	var (
		product pEntity.MstProduct
		err     error
	)
	product, err = s.productData.TampilDetailReceiveByNoReceive(ctx, NoTranrc)

	if err != nil {
		return product, errors.Wrap(err, "[SERVICE][TampilDetailReceiveByNoReceive]")
	}
	return product, err
}
