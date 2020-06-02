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
	GetAllHeaderReceive(ctx context.Context) ([]pEntity.HeaderRC, error)
	GetAllDetailReceive(ctx context.Context) ([]pEntity.DetailRC, error)
	GetDataHeaderByNoReceive(ctx context.Context, NoTranrc string) (pEntity.HeaderRC, error)
	GetDataDetailByNoReceive(ctx context.Context, NoTranrc string) ([]pEntity.DetailRC, error)
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

// TampilAllHeaderDataReceive ...
func (s Service) TampilAllHeaderDataReceive(ctx context.Context) ([]pEntity.HeaderRC, error) {
	var (
		header []pEntity.HeaderRC
		err    error
	)

	header, err = s.productData.GetAllHeaderReceive(ctx)
	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][TampilAllHeaderDataReceive")
	}

	return header, err
}

// TampilDataByNoReceive ...
func (s Service) TampilDataByNoReceive(ctx context.Context, NoTranrc string) (pEntity.JSONRCByNoReceive, error) {
	var (
		newDetails  []pEntity.DetailRC
		dataReceive pEntity.JSONRCByNoReceive
		produk      pEntity.MstProduct
		err         error
	)
	//Tampil Header
	dataReceive.HeaderRC, err = s.productData.GetDataHeaderByNoReceive(ctx, NoTranrc)
	if err != nil {
		return dataReceive, errors.Wrap(err, "[SERVICE][GetDataByNoReceiveHeader")
	}

	//Tampil Detail
	dataReceive.DetailRC, err = s.productData.GetDataDetailByNoReceive(ctx, NoTranrc)
	//Test print raw data yang diterima
	fmt.Println("receive : ", dataReceive.DetailRC)
	//Looping Insert Data from API
	for _, detail := range dataReceive.DetailRC {
		produk, err = s.mpData.GetAllJSONMP(ctx, detail.KodeProduct.String)

		if detail.KodeProduct == produk.ProCode {
			detail.DeskripsiProduct.SetValid(produk.ProName.String)
			detail.SellPack.SetValid(produk.ProSellPack.Int64)
			detail.ExpDate.SetValid(produk.ProExpDateYN.String)
		}

		newDetails = append(newDetails, detail)
		fmt.Println("Detail : ", detail)
		fmt.Println("newDetails : ", newDetails)
	}
	//Memasukan data baru ke dalam array Detail
	dataReceive.DetailRC = newDetails
	//Error Handling
	if err != nil {
		return dataReceive, errors.Wrap(err, "[SERVICE][GetDataByNoReceiveHeader")
	}

	//return Header dan Detail
	return dataReceive, err
}
