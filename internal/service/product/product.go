package product

import (
	"context"
	"fmt"
	doEntity "product/internal/entity/do"
	outletEntity "product/internal/entity/outlet"
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
	InsertDataHeaderFromAPI(ctx context.Context, header doEntity.TransfH) error
	InsertDataDetailFromAPI(ctx context.Context, detail doEntity.TransfD) error
	EditDetailOrderByNoTransfandProcode(ctx context.Context, noTransf string, procod string, detail doEntity.TransfD) (doEntity.TransfD, error)
	PrintReceive(ctx context.Context, noTransf string, NoTranrc string) ([]pEntity.JSONPrintReceive, error)
}

type mpData interface {
	GetAllJSONMP(ctx context.Context, kode string) (pEntity.MstProduct, error)
}

type outletData interface {
	GetOutletName(ctx context.Context, outcode string) (outletEntity.Outlet, error)
}

type doData interface {
	GetAllHeaderJSONDO(ctx context.Context, noTransf string) (doEntity.TransfH, error)
	GetAllDetailJSONDO(ctx context.Context, noTransf string) ([]doEntity.TransfD, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	productData Data
	mpData      mpData
	outletData  outletData
	doData      doData
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(productData Data, mpData mpData, outletData outletData, doData doData) Service {
	// Assign variable dari parameter ke object
	return Service{
		productData: productData,
		mpData:      mpData,
		outletData:  outletData,
		doData:      doData,
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
		headers    []pEntity.HeaderRC
		newHeaders []pEntity.HeaderRC
		outlet     outletEntity.Outlet
		err        error
	)

	headers, err = s.productData.GetAllHeaderReceive(ctx)
	//Looping Insert Data from API
	for _, header := range headers {
		outlet, err = s.outletData.GetOutletName(ctx, header.KodePengirim.String)

		if header.KodePengirim == outlet.OutCode {
			header.Pengirim.SetValid(outlet.OutName.String)
		}

		newHeaders = append(newHeaders, header)
	}
	//Memasukan data baru ke dalam array Header
	headers = newHeaders
	//Error Handling
	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][TampilAllHeaderDataReceive")
	}

	return headers, err
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

// TampilDataDO ...
func (s Service) TampilDataDO(ctx context.Context, noTransf string) (doEntity.JSONDO, error) {
	var (
		order doEntity.JSONDO
		err   error
	)

	order.Header, err = s.doData.GetAllHeaderJSONDO(ctx, noTransf)
	if err != nil {
		return order, errors.Wrap(err, "[SERVICE][TampilDataDOHeader")
	}
	fmt.Println("getHeader : ", order.Header)

	order.Detail, err = s.doData.GetAllDetailJSONDO(ctx, noTransf)
	if err != nil {
		return order, errors.Wrap(err, "[SERVICE][TampilDataDODetail")
	}

	return order, err
}

// InsertDataFromAPI ...
func (s Service) InsertDataFromAPI(ctx context.Context, noTransf string) (doEntity.JSONDO, error) {
	var (
		order doEntity.JSONDO
		err   error
	)
	//getHeader
	order.Header, err = s.doData.GetAllHeaderJSONDO(ctx, noTransf)
	if err != nil {
		return order, errors.Wrap(err, "[SERVICE][TampilDataHeaderFromAPI1")
	}
	fmt.Println("getHeader : ", order.Header)
	//insertHeader
	err = s.productData.InsertDataHeaderFromAPI(ctx, order.Header)
	if err != nil {
		return order, errors.Wrap(err, "[SERVICE][InsertDataHeaderFromAPI2")
	}

	//getDetail
	order.Detail, err = s.doData.GetAllDetailJSONDO(ctx, noTransf)
	if err != nil {
		return order, errors.Wrap(err, "[SERVICE][InsertDataDetailFromAPI1")
	}
	//insertDetail
	for _, detail := range order.Detail {
		//insertDetail
		err = s.productData.InsertDataDetailFromAPI(ctx, detail)
		if err != nil {
			return order, errors.Wrap(err, "[SERVICE][InsertDataDetailFromAPI2")
		}
	}
	return order, err
}

// EditDetailOrderByNoTransfandProcode ...
func (s Service) EditDetailOrderByNoTransfandProcode(ctx context.Context, noTransf string, procod string, detail doEntity.TransfD) error {
	var (
		err error
	)

	detail, err = s.productData.EditDetailOrderByNoTransfandProcode(ctx, noTransf, procod, detail)
	fmt.Println("data : ", detail)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return err
}

// PrintReceive ...
func (s Service) PrintReceive(ctx context.Context, noTransf string, NoTranrc string) ([]pEntity.JSONPrintReceive, error) {
	var (
		receives    []pEntity.JSONPrintReceive
		newReceives []pEntity.JSONPrintReceive
		produk      pEntity.MstProduct
		outlet      outletEntity.Outlet
		err         error
	)

	//Tampil Detail
	receives, err = s.productData.PrintReceive(ctx, noTransf, NoTranrc)
	//Test print raw data yang diterima
	fmt.Println("receive : ", receives)
	//Looping Insert Data from API
	for _, receive := range receives {
		produk, err = s.mpData.GetAllJSONMP(ctx, receive.KodeProduct.String)
		outlet, err = s.outletData.GetOutletName(ctx, receive.Pengirim.String)
		outlet, err = s.outletData.GetOutletName(ctx, receive.Penerima.String)

		if receive.KodeProduct == produk.ProCode {
			receive.DeskripsiProduct.SetValid(produk.ProName.String)
			receive.Satuan.SetValid(produk.ProSellPack.Int64)
		} else if receive.Pengirim == outlet.OutCode {
			receive.Pengirim.SetValid(outlet.OutName.String)
		} else if receive.Penerima == outlet.OutCode {
			receive.Penerima.SetValid(outlet.OutName.String)
		}

		newReceives = append(newReceives, receive)
		fmt.Println("Detail : ", receive)
		fmt.Println("newDetails : ", newReceives)
	}
	//Memasukan data baru ke dalam array Detail
	receives = newReceives
	//Error Handling
	if err != nil {
		return receives, errors.Wrap(err, "[SERVICE][PrintReceive")
	}

	//return Header dan Detail
	return receives, err
}

// // PrintReceive ...
// func (s Service) PrintReceive(ctx context.Context, noTransf string, NoTranrc string) (pEntity.JSONPrintReceive, error) {
// 	var (
// 		receive pEntity.JSONPrintReceive
// 		receives    []pEntity.JSONPrintReceive
// 		lists []pEntity.JSONintiPrintReceive
// 		newReceives []pEntity.JSONPrintReceive
// 		produk      pEntity.MstProduct
// 		outlet      outletEntity.Outlet
// 		err         error
// 	)

// 	//Tampil Detail
// 	receive, err = s.productData.PrintReceive(ctx, noTransf, NoTranrc)
// 	//Test print raw data yang diterima
// 	fmt.Println("receive : ", receives)
// 	//Looping Insert Data from API
// 	for _, receive := range receives {
// 		//produk, err = s.mpData.GetAllJSONMP(ctx, receive.KodeProduct.String)
// 		outlet, err = s.outletData.GetOutletName(ctx, receive.Pengirim.String)
// 		outlet, err = s.outletData.GetOutletName(ctx, receive.Penerima.String)

// 		//if receive.KodeProduct == produk.ProCode {
// 		//	receive.DeskripsiProduct.SetValid(produk.ProName.String)
// 		//	receive.Satuan.SetValid(produk.ProSellPack.Int64)
// 		//} else
// 		if receive.Pengirim == outlet.OutCode {
// 			receive.Pengirim.SetValid(outlet.OutName.String)
// 		} else if receive.Penerima == outlet.OutCode {
// 			receive.Penerima.SetValid(outlet.OutName.String)
// 		}

// 		newReceives = append(newReceives, receive)
// 		fmt.Println("Detail : ", receive)
// 		fmt.Println("newDetails : ", newReceives)
// 	}

// 	for _, list := range lists {
// 		produk, err = s.mpData.GetAllJSONMP(ctx, list.KodeProduct.String)
// 		if list.KodeProduct == produk.ProCode {
// 			list.DeskripsiProduct.SetValid(produk.ProName.String)
// 			list.Satuan.SetValid(produk.ProSellPack.Int64)
// 		}
// 	}
// 	//Memasukan data baru ke dalam array Detail
// 	receives = newReceives
// 	//Error Handling
// 	if err != nil {
// 		return receives, errors.Wrap(err, "[SERVICE][PrintReceive")
// 	}

// 	//return Header dan Detail
// 	return receives, err
// }
