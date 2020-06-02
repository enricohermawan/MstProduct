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
	GetAllDataMstProduct(ctx context.Context) (pEntity.MstProduct, error)
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

// InsertDataDetailFromAPI ...
func (s Service) InsertDataDetailFromAPI(ctx context.Context, NoTranrc string) ([]pEntity.DetailRC, error) {
	var (
		err        error
		details    []pEntity.DetailRC
		newDetail  pEntity.DetailRC
		newDetails []pEntity.DetailRC
		produk     pEntity.MstProduct
	)

	details, err = s.productData.GetDataDetailByNoReceive(ctx, NoTranrc)
	for _, detail := range details {
		produk, err = s.mpData.GetAllJSONMP(ctx, produk.ProCode.String)
		newDetail = detail
		fmt.Println("json : ", newDetail)
		if detail.KodeProduct == produk.ProCode {
			newDetail.DeskripsiProduct.SetValid(produk.ProName.String)
			newDetail.SellPack.SetValid(produk.ProSellPack.Int64)
			newDetail.ExpDate.SetValid(produk.ProExpDateYN.String)
		}

		if err != nil {
			return newDetails, errors.Wrap(err, "[SERVICE][InsertDataDetailFromAPI2")
		}
		newDetails = append(newDetails, newDetail)
	}

	return newDetails, err
}

// TampilAllDataReceive ...
func (s Service) TampilAllDataReceive(ctx context.Context) ([]pEntity.HeaderRC, error) {
	var (
		header []pEntity.HeaderRC
		err    error
	)

	header, err = s.productData.GetAllHeaderReceive(ctx)
	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetAllDataHeaderReceiveByNoTranrc")
	}

	// receive.DetailRC, err = s.productData.GetAllDetailReceive(ctx)
	// if err != nil {
	// 	return receive, errors.Wrap(err, "[SERVICE][GetAllDataDetailReceiveByNoTranrc")
	// }

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
	fmt.Println("receive : ", dataReceive.DetailRC)
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
	dataReceive.DetailRC = newDetails
	if err != nil {
		return dataReceive, errors.Wrap(err, "[SERVICE][GetDataByNoReceiveHeader")
	}

	//return Header dan Detail
	return dataReceive, err
}
