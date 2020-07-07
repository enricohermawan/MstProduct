package do

import (
	"time"

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
	TransfDLastUpdate      time.Time   `db:"" json:"TransfD_LastUpdate"`
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
	TransfHLastUpdate    time.Time   `db:"" json:"TransfH_LastUpdate"`
	TransfHDataAktifYN   zero.String `db:"" json:"TransfH_DataAktifYN"`
	TransfHOrderID       zero.String `db:"" json:"TransfH_OrderID"`
	TransfHSPID          zero.String `db:"" json:"TransfH_SPID"`
	TransfHPaymentMethod zero.String `db:"" json:"TransfH_PaymentMethod"`
}

// JSONDO ...
type JSONDO struct {
	// InsertTime zero.String `json:"insertTime"`
	Detail []TransfD `json:"transF_D"`
	Header TransfH   `json:"transf_h"`
}

// JSONDO1 ...
type JSONDO1 struct {
	Header TransfH `json:"transf_h"`
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

// JSONGenerateReceiveHeader ...
type JSONGenerateReceiveHeader struct {
	TransfHRunningID     zero.Int    `db:"" json:"TransfH_RunningID"`
	TransfHOutCodeTransf zero.String `db:"" json:"TransfH_OutCodeTransf"`
}

// InsertReceive ...
type InsertReceive struct {
	NoTranrc zero.String `db:"" json:"NoTranrc"`
}

// TransfDSaveGenerate ...
type TransfDSaveGenerate struct {
	TransfDRunningID       zero.Int    `db:"TransfD_RunningID" json:"TransfD_RunningID"`
	TransfDOutCodeTransf   zero.String `db:"TransfD_OutCodeTransf" json:"TransfD_OutCodeTransf"`
	TransfDNoTransf        zero.String `db:"TransfD_NoTransf" json:"TransfD_NoTransf"`
	NoTranrc               zero.String `db:"NoTranrc" json:"NoTranrc"`
	TransfDGroup           zero.Int    `db:"TransfD_Group" json:"TransfD_Group"`
	TransfDOutCodeSP       zero.String `db:"TransfD_OutCodeSP" json:"TransfD_OutCodeSP"`
	TransfDNoSP            zero.String `db:"TransfD_NoSP" json:"TransfD_NoSP"`
	TransfDProCod          zero.String `db:"TransfD_ProCod" json:"TransfD_ProCod"`
	TransfDBatchNumber     zero.String `db:"TransfD_BatchNumber" json:"TransfD_BatchNumber"`
	TransfDED              zero.String `db:"TransfD_ED" json:"TransfD_ED"`
	TransfDQty             zero.Int    `db:"TransfD_Qty" json:"TransfD_Qty"`
	TransfDQtyScan         zero.Int    `db:"TransfD_Qty_Scan" json:"TransfD_Qty_Scan"`
	TransfDQtyStk          zero.Int    `db:"TransfD_QtyStk" json:"TransfD_QtyStk"`
	TransfDOutCodeOrder    zero.String `db:"TransfD_OutCodeOrder" json:"TransfD_OutCodeOrder"`
	TransfDNoOrder         zero.String `db:"TransfD_NoOrder" json:"TransfD_NoOrder"`
	TransFDCategoryProduct zero.Int    `db:"TransFD_CategoryProduct" json:"TransFD_CategoryProduct"`
	TransFDEditYN          zero.String `db:"TransFD_EditYN" json:"TransFD_EditYN"`
	TransfDActiveYN        zero.String `db:"TransfD_ActiveYN" json:"TransfD_ActiveYN"`
	TransfDUserID          zero.String `db:"TransfD_UserID" json:"TransfD_UserID"`
	TransfDLastUpdate      time.Time   `db:"TransfD_LastUpdate" json:"TransfD_LastUpdate"`
	TransfDDataAktifYN     zero.String `db:"TransfD_DataAktifYN" json:"TransfD_DataAktifYN"`
	TransfDSPID            zero.String `db:"TransfD_SPID" json:"TransfD_SPID"`
	TransfDSalePrice       zero.Int    `db:"TransfD_SalePrice" json:"TransfD_SalePrice"`
	TransfDDiscount        zero.Int    `db:"TransfD_Discount" json:"TransfD_Discount"`
}

// TransfHSaveGenerate ...
type TransfHSaveGenerate struct {
	TransfHRunningID     zero.Int    `db:"TransfH_RunningID" json:"TransfH_RunningID"`
	TransfHOutCodeTransf zero.String `db:"TransfH_OutCodeTransf" json:"TransfH_OutCodeTransf"`
	TransfHNoTransf      zero.String `db:"TransfH_NoTransf" json:"TransfH_NoTransf"`
	TransfHTglTransf     zero.String `db:"TransfH_TglTransf" json:"TransfH_TglTransf"`
	NoTranrc             zero.String `db:"" json:"NoTranrc"`
	TransfHOutCodeDest   zero.String `db:"TransfH_OutCodeDest" json:"TransfH_OutCodeDest"`
	TransfHGroup         zero.Int    `db:"TransfH_Group" json:"TransfH_Group"`
	TransfHOutCodeSP     zero.String `db:"TransfH_OutCodeSP" json:"TransfH_OutCodeSP"`
	TransfHNoSP          zero.String `db:"TransfH_NoSP" json:"TransfH_NoSP"`
	TransfHTglSP         zero.String `db:"TransfH_TglSP" json:"TransfH_TglSP"`
	TransfHFlag          zero.String `db:"TransfH_Flag" json:"TransfH_Flag"`
	TransfHFlagTrf       zero.String `db:"TransfH_FlagTrf" json:"TransfH_FlagTrf"`
	TransfHTglDwld       zero.String `db:"TransfH_TglDwld" json:"TransfH_TglDwld"`
	TransfHFlagSttk      zero.Int    `db:"TransfH_FlagSttk" json:"TransfH_FlagSttk"`
	TransfHTranno        zero.Int    `db:"TransfH_Tranno" json:"TransfH_Tranno"`
	TransfHUpload        zero.String `db:"TransfH_Upload" json:"TransfH_Upload"`
	TransfHTransfer      zero.String `db:"TransfH_Transfer" json:"TransfH_Transfer"`
	TransFHLapakYN       zero.String `db:"TransFH_LapakYN" json:"TransFH_LapakYN"`
	TransfHActiveYN      zero.String `db:"TransfH_ActiveYN" json:"TransfH_ActiveYN"`
	TransfHUserID        zero.String `db:"TransfH_UserID" json:"TransfH_UserID"`
	TransfHLastUpdate    time.Time   `db:"TransfH_LastUpdate" json:"TransfH_LastUpdate"`
	TransfHDataAktifYN   zero.String `db:"TransfH_DataAktifYN" json:"TransfH_DataAktifYN"`
	TransfHOrderID       zero.String `db:"TransfH_OrderID" json:"TransfH_OrderID"`
	TransfHSPID          zero.String `db:"TransfH_SPID" json:"TransfH_SPID"`
	TransfHPaymentMethod zero.String `db:"TransfH_PaymentMethod" json:"TransfH_PaymentMethod"`
}

// JSONDOSaveGenerate ...
type JSONDOSaveGenerate struct {
	// InsertTime zero.String `json:"insertTime"`
	Detail []TransfDSaveGenerate `json:"transF_D"`
	Header TransfHSaveGenerate   `json:"transf_h"`
}
