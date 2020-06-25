package product

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

// MstProduct ...
type MstProduct struct {
	ProRunningID       zero.Int    `db:"Pro_RunningID" json:"pro_runningid"`
	ProCode            zero.String `db:"Pro_Code" json:"pro_code"`
	ProName            zero.String `db:"Pro_Name" json:"pro_name"`
	ProName2           zero.String `db:"Pro_Name2" json:"pro_name2"`
	ProDeptCode        zero.String `db:"Pro_DeptCode" json:"pro_deptcode"`
	ProPrinCode        zero.String `db:"Pro_PrinCode" json:"pro_princode"`
	ProCtrlCode        zero.String `db:"Pro_CtrlCode" json:"pro_ctrlcode"`
	ProStsMargin       zero.Int    `db:"Pro_StsMargin" json:"pro_stsmargin"`
	ProExpDateYN       zero.String `db:"Pro_ExpDateYN" json:"pro_expdateyn"`
	ProBuyPack         zero.Int    `db:"Pro_BuyPack" json:"pro_buypack"`
	ProSellUnit        zero.Int    `db:"Pro_SellUnit" json:"pro_sellunit"`
	ProSellPack        zero.Int    `db:"Pro_SellPack" json:"pro_sellpack"`
	ProMedUnit         zero.Int    `db:"Pro_MedUnit" json:"pro_medunit"`
	ProMedPack         zero.Int    `db:"Pro_MedPack" json:"pro_medpack"`
	ProBarcodeYN       zero.String `db:"Pro_BarcodeYN" json:"pro_barcodeyn"`
	ProMinOrdPO        zero.Float  `db:"Pro_MinOrdPO" json:"pro_minordpo"`
	ProKelipatan       zero.Float  `db:"Pro_Kelipatan" json:"pro_kelipatan"`
	ProNoRegYN         zero.String `db:"Pro_NoRegYN" json:"pro_noregyn"`
	ProNoReg           zero.String `db:"Pro_NoReg" json:"pro_noreg"`
	ProNIEDate         zero.String `db:"Pro_NIEDate" json:"pro_niedate"`
	ProKetNIE          zero.String `db:"Pro_KetNIE" json:"pro_ketnie"`
	ProNieSubmitBPOM   zero.String `db:"Pro_NieSubmitBPOM" json:"pro_niesubmitbpom"`
	ProHrgSpecialYN    zero.String `db:"Pro_HrgSpecialYN" json:"pro_hrgspecialyn"`
	ProHrgSpecial      zero.Float  `db:"Pro_HrgSpecial" json:"pro_hrgspecial"`
	ProHETYN           zero.String `db:"Pro_HETYN" json:"pro_hetyn"`
	ProHETPrice        zero.Float  `db:"Pro_HETPrice" json:"pro_hetprice"`
	ProBrandCode       zero.String `db:"Pro_BrandCode" json:"pro_brandcode"`
	ProNPDYN           zero.String `db:"Pro_NPDYN" json:"pro_npdyn"`
	ProTimbangYN       zero.String `db:"Pro_TimbangYN" json:"pro_timbangyn"`
	ProDecorYN         zero.String `db:"Pro_DecorYN" json:"pro_decoryn"`
	ProTabletYN        zero.String `db:"Pro_TabletYN" json:"pro_tabletyn"`
	ProMonograf        zero.String `db:"Pro_Monograf" json:"pro_monograf"`
	ProPatenYN         zero.String `db:"Pro_PatenYN" json:"pro_patenyn"`
	ProImportYN        zero.String `db:"Pro_ImportYN" json:"pro_importyn"`
	ProPackYN          zero.String `db:"Pro_PackYN" json:"pro_packyn"`
	ProKlasiBPOM       zero.String `db:"Pro_KlasiBPOM" json:"pro_klasibpom"`
	ProClass           zero.Int    `db:"Pro_Class" json:"pro_class"`
	ProIPA             zero.String `db:"Pro_IPA" json:"pro_ipa"`
	ProValueIPA        zero.Float  `db:"Pro_ValueIPA" json:"pro_valueipa"`
	ProIDA             zero.String `db:"Pro_IDA" json:"pro_ida"`
	ProValueIDA        zero.Float  `db:"Pro_ValueIDA" json:"pro_valueida"`
	ProGenerikID       zero.String `db:"Pro_GenerikID" json:"pro_generikid"`
	ProTerapiID        zero.String `db:"Pro_TerapiID" json:"pro_terapiid"`
	ProStrength        zero.Int    `db:"Pro_Strength" json:"pro_strength"`
	ProDsgForm         zero.Int    `db:"Pro_DsgForm" json:"pro_dsgform"`
	ProAmount          zero.Int    `db:"Pro_Amount" json:"pro_amount"`
	ProUnit            zero.Int    `db:"Pro_Unit" json:"pro_unit"`
	ProKdGenerik       zero.String `db:"Pro_KdGenerik" json:"pro_kdgenerik"`
	ProFirstActivity   zero.String `db:"Pro_FirstActivity" json:"pro_firstactivity"`
	ProSurveyYN        zero.String `db:"Pro_SurveyYN" json:"pro_surveyyn"`
	ProPrgDiscYN       zero.String `db:"Pro_PrgDiscYN" json:"pro_prgdiscyn"`
	ProCityCodeHJA     zero.Int    `db:"Pro_CityCodeHJA" json:"pro_citycodehja"`
	ProSalePrice       zero.Float  `db:"Pro_SalePrice" json:"pro_saleprice"`
	ProSalePriceBox    zero.Float  `db:"Pro_SalePriceBox" json:"pro_salepricebox"`
	ProSalePriceNon    zero.Float  `db:"Pro_SalePriceNon" json:"pro_salepricenon"`
	ProSalePriceNonBox zero.Float  `db:"Pro_SalePriceNonBox" json:"pro_salepricenonbox"`
	ProHPP             zero.Float  `db:"Pro_HPP" json:"pro_hpp"`
	ProScore           zero.Float  `db:"Pro_Score" json:"pro_score"`
	ProMarginComp      zero.Float  `db:"Pro_MarginComp" json:"pro_margincomp"`
	ProKeterangan      zero.Int    `db:"Pro_Keterangan" json:"pro_keterangan"`
	ProBuyGrp          zero.Float  `db:"Pro_BuyGrp" json:"pro_buygrp"`
	ProBuyGrpBox       zero.Float  `db:"Pro_BuyGrpBox" json:"pro_buygrpbox"`
	ProActiveYN        zero.String `db:"Pro_ActiveYN" json:"pro_activeyn"`
	ProUserID          zero.String `db:"Pro_UserID" json:"pro_userid"`
	ProLastUpdate      zero.String `db:"Pro_LastUpdate" json:"pro_lastupdate"`
	ProDataAktifYN     zero.String `db:"Pro_DataAktifYN" json:"pro_dataaktifyn"`
	ProHomeBrandYN     zero.String `db:"Pro_HomeBrandYN" json:"pro_homebrandyn"`
	ProLeadTime        zero.Int    `db:"Pro_LeadTime" json:"pro_leadtime"`
	ProBufferYN        zero.String `db:"Pro_BufferYN" json:"pro_bufferyn"`
	ProStorageTemp     zero.Int    `db:"Pro_StorageTemp" json:"pro_storagetemp"`
	ProClsProduct      zero.String `db:"Pro_ClsProduct" json:"pro_clsproduct"`
	ProClsMargin       zero.String `db:"Pro_ClsMargin" json:"pro_clsmargin"`
	ProInHealth        zero.String `db:"Pro_InHealth" json:"pro_inhealth"`
	ProCodeSup         zero.String `db:"Pro_CodeSup" json:"pro_codesup"`
	ProPointMedPack    zero.Float  `db:"Pro_PointMedPack" json:"pro_pointmedpack"`
	ProPointSellPack   zero.Float  `db:"Pro_PointSellPack" json:"pro_pointsellpack"`
	ProPointDyn        zero.Float  `db:"Pro_PointDyn" json:"pro_pointdyn"`
	ProPointDynX       zero.Float  `db:"Pro_PointDynX" json:"pro_pointdynx"`
	ProPointDynO       zero.Float  `db:"Pro_PointDynO" json:"pro_pointdyno"`
	ProPointDynV       zero.Float  `db:"Pro_PointDynV" json:"pro_pointdynv"`
	ProBundleYN        zero.String `db:"Pro_BundleYN" json:"pro_bundleyn"`
	ProBolehBundleYN   zero.String `db:"Pro_BolehBundleYN" json:"pro_bolehbundleyn"`
	ProNonGenerikYN    zero.String `db:"Pro_NonGenerikYN" json:"pro_nongenerikyn"`
	ProNomorIjinPro    zero.String `db:"Pro_NomorIjinPro" json:"pro_nomorijinpro"`
	ProKategoriPro     zero.String `db:"Pro_KategoriPro" json:"pro_kategoripro"`
	ProHalalYN         zero.String `db:"Pro_HalalYN" json:"pro_halalyn"`
	ProNoSertHalal     zero.String `db:"Pro_NoSertHalal" json:"pro_noserthalal"`
	ProHalalDate       zero.String `db:"Pro_HalalDate" json:"pro_halaldate"`
	ProHalal           zero.String `db:"Pro_Halal" json:"pro_halal"`
}

// JSONTerima ...
type JSONTerima struct {
	Data      MstProduct   `json:"data"`
	DataArray []MstProduct `json:"data_array"`
}

// TransFD ...
type TransFD struct {
	TransfDRunningID       int    `db:"TransfD_RunningID" json:"TransfD_RunningID,omitempty"`
	TransfDOutCodeTransf   string `db:"TransfD_OutCodeTransf" json:"TransfD_OutCodeTransf,omitempty"`
	TransfDNoTransf        string `db:"TransfD_NoTransf" json:"TransfD_NoTransf,omitempty"`
	TransfDGroup           string `db:"TransfD_Group" json:"TransfD_Group,omitempty"`
	TransfDOutCodeSP       string `db:"TransfD_OutCodeSP" json:"TransfD_OutCodeSP,omitempty"`
	TransfDNoSP            string `db:"TransfD_NoSP" json:"TransfD_NoSP,omitempty"`
	TransfDProCod          string `db:"TransfD_ProCod" json:"TransfD_ProCod,omitempty"`
	TransfDBatchNumber     string `db:"TransfD_BatchNumber" json:"TransfD_BatchNumber,omitempty"`
	TransfDED              string `db:"TransfD_ED" json:"TransfD_ED,omitempty"`
	TransfDQty             int    `db:"TransfD_Qty" json:"TransfD_Qty,omitempty"`
	TransfDQtyScan         int    `db:"TransfD_Qty_Scan" json:"TransfD_Qty_Scan,omitempty"`
	TransfDQtyStk          int    `db:"TransfD_QtyStk" json:"TransfD_QtyStk,omitempty"`
	TransfDOutCodeOrder    string `db:"TransfD_OutCodeOrder" json:"TransfD_OutCodeOrder,omitempty"`
	TransfDNoOrder         string `db:"TransfD_NoOrder" json:"TransfD_NoOrder,omitempty"`
	TransFDCategoryProduct int    `db:"TransFD_CategoryProduct" json:"TransFD_CategoryProduct,omitempty"`
	TransFDEditYN          string `db:"TransFD_EditYN" json:"TransFD_EditYN,omitempty"`
	TransfDActiveYN        string `db:"TransfD_ActiveYN" json:"TransfD_ActiveYN,omitempty"`
	TransfDUserID          string `db:"TransfD_UserID" json:"TransfD_UserID,omitempty"`
	TransfDLastUpdate      string `db:"TransfD_LastUpdate" json:"TransfD_LastUpdate,omitempty"`
	TransfDDataAktifYN     string `db:"TransfD_DataAktifYN" json:"TransfD_DataAktifYN,omitempty"`
	TransfDSPID            string `db:"TransfD_SPID" json:"TransfD_SPID,omitempty"`
	TransfDSalePrice       string `db:"TransfD_SalePrice" json:"TransfD_SalePrice,omitempty"`
	TransfDDiscount        string `db:"TransfD_Discount" json:"TransfD_Discount,omitempty"`
}

// TransFH ...
type TransFH struct {
	TransfHNoSP          zero.String `db:"" json:"transfH_NoSP"`
	TransfHRunningID     zero.Int    `db:"" json:"transfH_RunningID"`
	TransfHLastUpdate    time.Time   `db:"" json:"transfH_LastUpdate"`
	TransfHUpload        zero.String `db:"" json:"transfH_Upload"`
	TransfHGroup         zero.Int    `db:"" json:"transfH_Group"`
	TransfHFlag          zero.String `db:"" json:"transfH_Flag"`
	TransfHDataAktifYN   zero.String `db:"" json:"transfH_DataAktifYN"`
	TransfHNoTransf      zero.String `db:"" json:"transfH_NoTransf"`
	TransfHTglTransf     time.Time   `db:"" json:"transfH_TglTransf"`
	TransfHOutCodeSP     zero.Int    `db:"" json:"transfH_OutCodeSP"`
	TransfHOutCodeDest   zero.Int    `db:"" json:"transfH_OutCodeDest"`
	TransfHOutCodeTransf zero.Int    `db:"" json:"transfH_OutCodeTransf"`
	TransfHFlagSttk      zero.String `db:"" json:"transfH_FlagSttk"`
	TransfHTransfer      zero.String `db:"" json:"transfH_Transfer"`
	TransfHActiveYN      zero.String `db:"" json:"transfH_ActiveYN"`
	TransfHTglDwld       time.Time   `db:"" json:"transfH_TglDwld"`
	TransfHUserID        zero.String `db:"" json:"transfH_UserID"`
	TransfHFlagTrf       zero.String `db:"" json:"transfH_FlagTrf"`
	TransfHTranno        zero.Int    `db:"" json:"transfH_Tranno"`
}

// TranRCD ...
type TranRCD struct {
	TranrcDRunningID     string `db:"TranrcD_RunningID" json:"TranrcD_RunningID"`
	TranrcDOutCodeTransf string `db:"TranrcD_OutCodeTransf" json:"TranrcD_OutCodeTransf"`
	TranrcDNoTransf      string `db:"TranrcD_NoTransf" json:"TranrcD_NoTransf"`
	TranrcDOutCodeTranrc string `db:"TranrcD_OutCodeTranrc" json:"TranrcD_OutCodeTranrc"`
	TranrcDNoTranrc      string `db:"TranrcD_NoTranrc" json:"TranrcD_NoTranrc"`
	TranrcDProcod        string `db:"TranrcD_Procod" json:"TranrcD_Procod"`
	TranrcDBBatchNumber  string `db:"TranrcDB_BatchNumber" json:"TranrcDB_BatchNumber"`
	TranrcDBKonsentrasi  string `db:"TranrcDB_Konsentrasi" json:"TranrcDB_Konsentrasi"`
	TranrcDQuantityRecv  string `db:"TranrcD_QuantityRecv" json:"TranrcD_QuantityRecv"`
	TranrcDQuantityScan  string `db:"TranrcD_QuantityScan" json:"TranrcD_QuantityScan"`
	TranrcDQuantityStk   string `db:"TranrcD_QuantityStk" json:"TranrcD_QuantityStk"`
	TranrcDOutCodeOR     string `db:"TranrcD_OutCodeOR" json:"TranrcD_OutCodeOR"`
	TranrcDNoOR          string `db:"TranrcD_NoOR" json:"TranrcD_NoOR"`
	TranrcDQuantityOR    string `db:"TranrcD_QuantityOR" json:"TranrcD_QuantityOR"`
	TranrcDActiveYN      string `db:"TranrcD_ActiveYN" json:"TranrcD_ActiveYN"`
	TranrcDUserID        string `db:"TranrcD_UserId" json:"TranrcD_UserId"`
	TranrcDlastUpdate    string `db:"TranrcD_lastUpdate" json:"TranrcD_lastUpdate"`
	TranrcDDataAktifYN   string `db:"TranrcD_DataAktifYN" json:"TranrcD_DataAktifYN"`
}

// TranRCH ...
type TranRCH struct {
	TranrcHRunningID     string    `db:"TranrcH_RunningID" json:"TranrcH_RunningID"`
	TranrcHOutCodeTransf string    `db:"TranrcH_OutCodeTransf" json:"TranrcH_OutCodeTransf"`
	TranrcHNoTransf      string    `db:"TranrcH_NoTransf" json:"TranrcH_NoTransf"`
	TranrcHOutCodeTranrc string    `db:"TranrcH_OutCodeTranrc" json:"TranrcH_OutCodeTranrc"`
	TranrcHNoTranrc      string    `db:"TranrcH_NoTranrc" json:"TranrcH_NoTranrc"`
	TranrcHTglTranrc     time.Time `db:"TranrcH_TglTranrc" json:"TranrcH_TglTranrc"`
	TranrcHFlag          string    `db:"TranrcH_Flag" json:"TranrcH_Flag"`
	TranrcHFlagTrf       string    `db:"TranrcH_FlagTrf" json:"TranrcH_FlagTrf"`
	TranrcHTglDwld       string    `db:"TranrcH_TglDwld" json:"TranrcH_TglDwld"`
	TranrcHNip           string    `db:"TranrcH_Nip" json:"TranrcH_Nip"`
	TranrcHActiveYN      string    `db:"TranrcH_ActiveYN" json:"TranrcH_ActiveYN"`
	TranrcHUserID        string    `db:"TranrcH_UserId" json:"TranrcH_UserId"`
	TranrcHlastUpdate    time.Time `db:"TranrcH_lastUpdate" json:"TranrcH_lastUpdate"`
	TranrcHDataAktifYN   string    `db:"TranrcH_DataAktifYN" json:"TranrcH_DataAktifYN"`
}

// GabunganMStProdDanTranRCD ...
type GabunganMStProdDanTranRCD struct {
	MstProduct
	TranRCD
}

// DetailRC ...
type DetailRC struct {
	KodeProduct      zero.String `db:"TranrcD_Procod" json:"KodeProduct"`
	DeskripsiProduct zero.String `json:"DeskripsiProduct"`
	QuantityScan     zero.Int    `db:"TranrcD_QuantityScan" json:"QuantityScan"`
	SellPack         zero.Int    `json:"SellPack"`
	QuantityDO       zero.Int    `db:"TranrcD_QuantityRecv" json:"QuantityDO"`
	QuantityStok     zero.String `db:"TranrcD_QuantityStk" json:"QuantityStok"`
	ExpDate          zero.String `json:"ExpDate"`
	BatchNumber      zero.String `db:"TranrcDB_BatchNumber" json:"BatchNumber"`
}

// HeaderRC ...
type HeaderRC struct {
	NomorReceive   zero.String `db:"TranrcH_NoTranrc" json:"NomorReceive"`
	TanggalReceive time.Time   `db:"TranrcH_TglTranrc" json:"TanggalReceive"`
	KodePengirim   zero.String `db:"TranrcH_OutCodeTransf" json:"KodePengirim"`
	Pengirim       zero.String `json:"Pengirim"`
	NomorTransfer  zero.String `db:"TranrcH_NoTransf" json:"NomorTransfer"`
	StatusPrint    zero.String `db:"TranrcH_Flag" json:"StatusPrint"`
}

// JSONRC ...
type JSONRC struct {
	HeaderRC []HeaderRC
	DetailRC []DetailRC
}

//JSONRCByNoReceive ...
type JSONRCByNoReceive struct {
	HeaderRC HeaderRC
	DetailRC []DetailRC
}

// HeaderDanTranRCH ...
type HeaderDanTranRCH struct {
	TranRCH
	HeaderRC
}

// DetailDanTranRCD ...
type DetailDanTranRCD struct {
	TranRCD
	DetailRC
}

// TransfD ...
type TransfD struct {
	TransfDRunningID     string `db:"TransfD_RunningID" json:"TransfD_RunningID"`
	TransfDOutCodeTransf string `db:"TransfD_OutCodeTransf" json:"TransfD_OutCodeTransf"`
	TransfDNoTransf      string `db:"TransfD_NoTransf" json:"TransfD_NoTransf"`
	TransfDOutCodeTranrc string `db:"TransfD_OutCodeTranrc" json:"TransfD_OutCodeTranrc"`
	TransfDDNoTranrc     string `db:"TransfD_NoTranrc" json:"TransfD_NoTranrc"`
	TransfDProcod        string `db:"TransfD_Procod" json:"TransfD_Procod"`
	TransfDBBatchNumber  string `db:"TransfD_BatchNumber" json:"TransfD_BatchNumber"`
	TransfDBKonsentrasi  string `db:"TransfD_Konsentrasi" json:"TransfD_Konsentrasi"`
	TransfDQuantityRecv  string `db:"TransfD_QuantityRecv" json:"TransfD_QuantityRecv"`
	TransfDQuantityScan  string `db:"TransfD_QuantityScan" json:"TransfD_QuantityScan"`
	TransfDQuantityStk   string `db:"TransfD_QuantityStk" json:"TransfD_QuantityStk"`
	TransfDOutCodeOR     string `db:"TransfD_OutCodeOR" json:"TransfD_OutCodeOR"`
	TransfDNoOR          string `db:"TransfD_NoOR" json:"TransfD_NoOR"`
	TransfDQuantityOR    string `db:"TransfD_QuantityOR" json:"TransfD_QuantityOR"`
	TransfDActiveYN      string `db:"TransfD_ActiveYN" json:"TransfD_ActiveYN"`
	TransfDUserID        string `db:"TransfD_UserId" json:"TransfD_UserId"`
	TransfDlastUpdate    string `db:"TransfD_lastUpdate" json:"TransfD_lastUpdate"`
	TransfDDataAktifYN   string `db:"TransfD_DataAktifYN" json:"TransfD_DataAktifYN"`
}

// TransfH ...
type TransfH struct {
	TransfHRunningID     string    `db:"TranrsfH_RunningID" json:"TranrsfH_RunningID"`
	TransfHOutCodeTransf string    `db:"TranrsfH_OutCodeTransf" json:"TranrsfH_OutCodeTransf"`
	TransfHNoTransf      string    `db:"TranrsfH_NoTransf" json:"TranrsfH_NoTransf"`
	TransfHOutCodeTranrc string    `db:"TranrsfH_OutCodeTranrc" json:"TranrsfH_OutCodeTranrc"`
	TransfHNoTranrc      string    `db:"TransfH_NoTranrc" json:"TransfH_NoTranrc"`
	TransfHTglTranrc     time.Time `db:"TransfH_TglTranrc" json:"TransfH_TglTranrc"`
	TransfHFlag          string    `db:"TransfH_Flag" json:"TransfH_Flag"`
	TransfHFlagTrf       string    `db:"TransfH_FlagTrf" json:"TransfH_FlagTrf"`
	TransfHTglDwld       string    `db:"TransfH_TglDwld" json:"TransfH_TglDwld"`
	TransfHNip           string    `db:"TransfH_Nip" json:"TransfH_Nip"`
	TransfHActiveYN      string    `db:"TransfH_ActiveYN" json:"TransfH_ActiveYN"`
	TransfHUserID        string    `db:"TransfH_UserId" json:"TransfH_UserId"`
	TransfHlastUpdate    time.Time `db:"TransfH_lastUpdate" json:"TransfH_lastUpdate"`
	TransfHDataAktifYN   string    `db:"TransfH_DataAktifYN" json:"TransfH_DataAktifYN"`
}

// // JSONintiPrintReceive ...
// type JSONintiPrintReceive struct {
// 	KodeProduct      zero.String `db:"TranrcD_Procod" json:"KodeProduct"`
// 	DeskripsiProduct zero.String `db:"" json:"DeskripsiProduct"`
// 	Jumlah           zero.Int    `db:"TranrcD_QuantityScan" json:"Jumlah"`
// 	Satuan           zero.Int    `db:"" json:"Satuan"`
// }

// // JSONheaderPrintReceive ...
// type JSONheaderPrintReceive struct {
// 	NomorTransfer  zero.String `db:"TranrcH_NoTransf" json:"NomorTransfer"`
// 	NomorReceive   zero.String `db:"TranrcH_NoTranrc" json:"NomorReceive"`
// 	TangaalReceive time.Time   `db:"TranrcH_TglTranrc" json:"TangaalReceive"`
// 	Pengirim       zero.String `db:"TranrcH_OutCodeTransf" json:"Pengirim"`
// 	Penerima       zero.String `db:"TranrcH_OutCodeTranrc" json:"Penerima"`
// }

// JSONPrintReceive ...
type JSONPrintReceive struct {
	NomorTransfer    zero.String `db:"TranrcH_NoTransf" json:"NomorTransfer"`
	NomorReceive     zero.String `db:"TranrcH_NoTranrc" json:"NomorReceive"`
	TangaalReceive   time.Time   `db:"TranrcH_TglTranrc" json:"TangaalReceive"`
	Pengirim         zero.String `db:"TranrcH_OutCodeTransf" json:"Pengirim"`
	Penerima         zero.String `db:"TranrcH_OutCodeTranrc" json:"Penerima"`
	KodeProduct      zero.String `db:"TranrcD_Procod" json:"KodeProduct"`
	DeskripsiProduct zero.String `db:"" json:"DeskripsiProduct"`
	Jumlah           zero.Int    `db:"TranrcD_QuantityScan" json:"Jumlah"`
	Satuan           zero.Int    `db:"" json:"Satuan"`
}

// // JSONPrintReceive ...
// type JSONPrintReceive struct {
// 	JSONHeader JSONheaderPrintReceive
// 	JSONDetail []JSONintiPrintReceive
// }
