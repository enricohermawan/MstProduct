package do

import (
	"gopkg.in/guregu/null.v3/zero"
)

// TransfD ...
type TransfD struct {
	TransfDRunningID       zero.Int    `db:"" json:"TransfD_RunningID"`
	TransfDOutCodeTransf   zero.String `db:"" json:"TransfD_OutCodeTransf"`
	TransfDNoTransf        zero.String `db:"" json:"TransfD_NoTransf"`
	TransfDGroup           zero.Int    `db:"" json:"TransfD_Group"`
	TransfDOutCodeSP       zero.String `db:"" json:"TransfD_OutCodeSP"`
	TransfDNoSP            zero.String `db:"" json:"TransfD_NoSP"`
	TransfDProCod          zero.String `db:"" json:"TransfD_ProCod"`
	TransfDBatchNumber     zero.String `db:"" json:"TransfD_BatchNumber"`
	TransfDED              zero.String `db:"" json:"TransfD_ED"`
	TransfDQty             zero.Int    `db:"" json:"TransfD_Qty"`
	TransfDQtyScan         zero.Int    `db:"" json:"TransfD_QtyScan"`
	TransfDQtyStk          zero.Int    `db:"" json:"TransfD_QtyStk"`
	TransfDOutCodeOrder    zero.String `db:"" json:"TransfD_OutCodeOrder"`
	TransfDNoOrder         zero.String `db:"" json:"TransfD_NoOrder"`
	TransFDCategoryProduct zero.Int    `db:"" json:"TransFD_CategoryProduct"`
	TransFDEditYN          zero.String `db:"" json:"TransFD_EditYN"`
	TransfDActiveYN        zero.String `db:"" json:"TransfD_ActiveYN"`
	TransfDUserID          zero.String `db:"" json:"TransfD_UserID"`
	TransfDLastUpdate      zero.String `db:"" json:"TransfD_LastUpdate"`
	TransfDDataAktifYN     zero.String `db:"" json:"TransfD_DataAktifYN"`
	TransfDSPID            zero.String `db:"" json:"TransfD_SPID"`
	TransfDSalePrice       zero.Int    `db:"" json:"TransfD_SalePrice"`
	TransfDDiscount        zero.Int    `db:"" json:"TransfD_Discount"`
}

// TransfH ...
type TransfH struct {
	TransfHRunningID     zero.Int    `db:"" json:"TransfH_RunningID"`
	TransfHOutCodeTransf zero.String `db:"" json:"TransfH_OutCodeTransf"`
	TransfHNoTransf      zero.String `db:"" json:"TransfH_NoTransf"`
	TransfHTglTransf     zero.String `db:"" json:"TransfH_TglTransf"`
	TransfHOutCodeDest   zero.String `db:"" json:"TransfH_OutCodeDest"`
	TransfHGroup         zero.Int    `db:"" json:"TransfH_Group"`
	TransfHOutCodeSP     zero.String `db:"" json:"TransfH_OutCodeSP"`
	TransfHNoSP          zero.String `db:"" json:"TransfH_NoSP"`
	TransfHTglSP         zero.String `db:"" json:"TransfH_TglSP"`
	TransfHFlag          zero.String `db:"" json:"TransfH_Flag"`
	TransfHFlagTrf       zero.String `db:"" json:"TransfH_FlagTrf"`
	TransfHTglDwld       zero.String `db:"" json:"TransfH_TglDwld"`
	TransfHFlagSttk      zero.Int    `db:"" json:"TransfH_FlagSttk"`
	TransfHTranno        zero.Int    `db:"" json:"TransfH_Tranno"`
	TransfHUpload        zero.String `db:"" json:"TransfH_Upload"`
	TransfHTransfer      zero.String `db:"" json:"TransfH_Transfer"`
	TransFHLapakYN       zero.String `db:"" json:"TransFH_LapakYN"`
	TransfHActiveYN      zero.String `db:"" json:"TransfH_ActiveYN"`
	TransfHUserID        zero.String `db:"" json:"TransfH_UserID"`
	TransfHLastUpdate    zero.String `db:"" json:"TransfH_LastUpdate"`
	TransfHDataAktifYN   zero.String `db:"" json:"TransfH_DataAktifYN"`
	TransfHOrderID       zero.String `db:"" json:"TransfH_OrderID"`
	TransfHSPID          zero.String `db:"" json:"TransfH_SPID"`
	TransfHPaymentMethod zero.String `db:"" json:"TransfH_PaymentMethod"`
}

// JSONDO ...
type JSONDO struct {
	Header TransfH
	Detail []TransfD
}

// // DetailRC ...
// type DetailRC struct {
// 	KodeProduct      zero.String `db:"" json:"KodeProduct"`
// 	DeskripsiProduct zero.String `json:"DeskripsiProduct"`
// 	QuantityScan     zero.Int    `db:"" json:"QuantityScan"`
// 	SellPack         zero.Int    `json:"SellPack"`
// 	QuantityDO       zero.Int    `db:"" json:"QuantityDO"`
// 	QuantityStok     zero.String `db:"" json:"QuantityStok"`
// 	ExpDate          zero.String `json:"ExpDate"`
// 	BatchNumber      zero.String `db:"" json:"BatchNumber"`
// }

// // HeaderRC ...
// type HeaderRC struct {
// 	NomorReceive   zero.String `db:"" json:"NomorReceive"`
// 	TanggalReceive time.Time   `db:"" json:"TanggalReceive"`
// 	KodePengirim   zero.String `db:"" json:"KodePengirim"`
// 	Pengirim       zero.String `json:"Pengirim"`
// 	NomorTransfer  zero.String `db:"" json:"NomorTransfer"`
// 	StatusPrint    zero.String `db:"" json:"StatusPrint"`
// }

// //JSONTampilanByNoTransf ...
// type JSONTampilanByNoTransf struct {
// 	HeaderRC HeaderRC
// 	DetailRC []DetailRC
// }

// EditProduct ...
type EditProduct struct {
	Detail TransfD `json:"detail"`
}
