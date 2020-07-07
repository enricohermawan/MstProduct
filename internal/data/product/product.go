package product

import (
	"context"
	"log"
	doEntity "product/internal/entity/do"
	pEntity "product/internal/entity/product"
	"product/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	getDataHeaderByNoReceive  = "GetDataHeaderByNoReceive"
	qgetDataHeaderByNoReceive = `SELECT TranrcH_NoTranrc,TranrcH_TglTranrc,TranrcH_OutCodeTransf,
	TranrcH_NoTransf,TranrcH_Flag FROM TranRCH WHERE TranrcH_NoTranrc = ?`

	getDataDetailByNoReceive  = "GetDataDetailByNoReceive"
	qgetDataDetailByNoReceive = `SELECT TranrcD_Procod,TranrcD_QuantityScan,TranrcD_QuantityRecv,TranrcD_QuantityStk,
	TranrcDB_BatchNumber FROM TranRCD WHERE TranrcD_NoTranrc = ?`

	getAllDetailReceive  = "GetAllDetailReceive"
	qgetAllDetailReceive = `SELECT TranrcD_Procod,TranrcD_QuantityScan,TranrcD_QuantityRecv,TranrcD_QuantityStk,
	TranrcDB_BatchNumber FROM TranRCD `

	getAllHeaderReceive  = "GetAllHeaderReceive"
	qgetAllHeaderReceive = `SELECT TranrcH_NoTranrc,TranrcH_TglTranrc,TranrcH_OutCodeTransf,
	TranrcH_NoTransf,TranrcH_Flag FROM TranRCH WHERE TranrcH_DataAktifYN = 'Y'`

	insertDataHeaderFromAPI  = "InsertDataHeaderFromAPI"
	qinsertDataHeaderFromAPI = "INSERT INTO TransfH VALUES (NULL,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	insertDataDetailFromAPI  = "InsertDataDetailFromAPI"
	qinsertDataDetailFromAPI = "INSERT INTO TransfD VALUES (NULL,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	editDataDetail  = "EditDataDetail"
	qEditDataDetail = `UPDATE TransfD
	SET
	TransfD_BatchNumber = ?,
	TransfD_Qty_Scan = ?
	WHERE TransfD_NoTransf = ? AND TransfD_Procod = ?`

	printReceive  = "PrintReceive"
	qprintReceive = `SELECT TranrcH_NoTransf,TranrcH_NoTranrc,TranrcH_TglTranrc,TranrcH_OutCodeTransf,
	TranrcH_OutCodeTranrc,TranrcD_Procod,TranrcD_QuantityScan FROM TranRCD JOIN TranRCH ON TranRCD.TranrcD_NoTransf
	= TranRCH.TranrcH_NoTransf WHERE TranrcH_NoTransf = ? AND TranrcH_NoTranrc = ?`

	generateReceiveHeader  = "GenerateReceiveHeader"
	qgenerateReceiveHeader = `SELECT 
    CASE 
    WHEN LENGTH(MAX(RIGHT(TranrcH_NoTranrc,4)+1)) = 1
    THEN CONCAT('R',DATE_FORMAT(NOW(), '%y%m'),'000',MAX(RIGHT(TranrcH_NoTranrc,4)+1)) 

    WHEN LENGTH(MAX(RIGHT(TranrcH_NoTranrc,4)+1)) = 2
    THEN CONCAT('R',DATE_FORMAT(NOW(), '%y%m'),'00',MAX(RIGHT(TranrcH_NoTranrc,4)+1)) 

    WHEN LENGTH(MAX(RIGHT(TranrcH_NoTranrc,4)+1)) = 3
    THEN CONCAT('R',DATE_FORMAT(NOW(), '%y%m'),'0',MAX(RIGHT(TranrcH_NoTranrc,4)+1)) 

    ELSE CONCAT('R',DATE_FORMAT(NOW(), '%y%m'),MAX(RIGHT(TranrcH_NoTranrc,4)+1)) 
    END AS TranrcH_NoTranrc
	FROM TranRCH`

	saveAndGenerateTranRCH  = "SaveAndGenerateTranRCH"
	qsaveAndGenerateTranRCH = "INSERT INTO TranRCH VALUES (NULL,?,?,?,?,?,?,?,?,NULL,?,NULL,?,?)"

	saveAndGenerateTranRCD  = "SaveAndGenerateTranRCD"
	qsaveAndGenerateTranRCD = "INSERT INTO TranRCD VALUES (NULL,?,?,?,?,?,?,NULL,?,?,?,NULL,NULL,NULL,?,NULL,?,?)"

	insertReceive  = "InsertReceive"
	qinsertReceive = `UPDATE TranRCH 
	SET 
	TranrcH_NoTranrc = ?
	WHERE TranrcH_NoTransf = ?`

	getDataDOTransfH  = "GetDataDOTransfH"
	qgetDataDOTransfH = "SELECT * FROM TransfH WHERE TransfH_NoTransf = ?"

	getDataDOTransfD  = "GetDataDOTransfD"
	qgetDataDOTransfD = "SELECT * FROM TransfD WHERE TransfD_NoTransf = ?"
)

var (
	readStmt = []statement{

		{getAllDetailReceive, qgetAllDetailReceive},
		{getAllHeaderReceive, qgetAllHeaderReceive},
		{getDataDetailByNoReceive, qgetDataDetailByNoReceive},
		{getDataHeaderByNoReceive, qgetDataHeaderByNoReceive},
		{insertDataHeaderFromAPI, qinsertDataHeaderFromAPI},
		{insertDataDetailFromAPI, qinsertDataDetailFromAPI},
		{editDataDetail, qEditDataDetail},
		{printReceive, qprintReceive},
		{generateReceiveHeader, qgenerateReceiveHeader},
		{saveAndGenerateTranRCH, qsaveAndGenerateTranRCH},
		{saveAndGenerateTranRCD, qsaveAndGenerateTranRCD},
		{insertReceive, qinsertReceive},
		{getDataDOTransfH, qgetDataDOTransfH},
		{getDataDOTransfD, qgetDataDOTransfD},
	}
)

// New ...
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetAllHeaderReceive ...
func (d Data) GetAllHeaderReceive(ctx context.Context) ([]pEntity.HeaderRC, error) {
	var (
		rows    *sqlx.Rows
		header  pEntity.HeaderRC
		headers []pEntity.HeaderRC
		err     error
	)
	rows, err = d.stmt[getAllHeaderReceive].QueryxContext(ctx)
	for rows.Next() {
		if err := rows.StructScan(&header); err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllDetailReceive] ")
		}
		headers = append(headers, header)
	}

	return headers, err
}

// GetAllDetailReceive ...
func (d Data) GetAllDetailReceive(ctx context.Context) ([]pEntity.DetailRC, error) {
	var (
		rows    *sqlx.Rows
		detail  pEntity.DetailRC
		details []pEntity.DetailRC
		err     error
	)
	rows, err = d.stmt[getAllDetailReceive].QueryxContext(ctx)
	for rows.Next() {
		if err := rows.StructScan(&detail); err != nil {
			return details, errors.Wrap(err, "[DATA][GetAllDetailReceive] ")
		}
		details = append(details, detail)
	}

	return details, err
}

// GetDataHeaderByNoReceive ...
func (d Data) GetDataHeaderByNoReceive(ctx context.Context, NoTranrc string) (pEntity.HeaderRC, error) {
	var (
		header pEntity.HeaderRC
		err    error
	)

	if err := d.stmt[getDataHeaderByNoReceive].QueryRowxContext(ctx, NoTranrc).StructScan(&header); err != nil {
		return header, errors.Wrap(err, "[DATA][GetDataHeaderByNoReceive]")
	}

	return header, err
}

// GetDataDetailByNoReceive ...
func (d Data) GetDataDetailByNoReceive(ctx context.Context, NoTranrc string) ([]pEntity.DetailRC, error) {
	var (
		rows    *sqlx.Rows
		detail  pEntity.DetailRC
		details []pEntity.DetailRC
		err     error
	)
	rows, err = d.stmt[getDataDetailByNoReceive].QueryxContext(ctx, NoTranrc)
	for rows.Next() {
		if err := rows.StructScan(&detail); err != nil {
			return details, errors.Wrap(err, "[DATA][GetDataDetailByNoReceive] ")
		}
		details = append(details, detail)
	}

	return details, err
}

// InsertDataHeaderFromAPI ...
func (d Data) InsertDataHeaderFromAPI(ctx context.Context, header doEntity.TransfH) error {
	var err error
	if _, err = d.stmt[insertDataHeaderFromAPI].ExecContext(ctx,
		header.TransfHOutCodeTransf,
		header.TransfHNoTransf,
		header.TransfHTglTransf,
		header.TransfHOutCodeDest,
		header.TransfHGroup,
		header.TransfHOutCodeSP,
		header.TransfHNoSP,
		header.TransfHTglSP,
		header.TransfHFlag,
		header.TransfHFlagTrf,
		header.TransfHTglDwld,
		header.TransfHFlagSttk,
		header.TransfHTranno,
		header.TransfHUpload,
		header.TransfHTransfer,
		header.TransFHLapakYN,
		header.TransfHActiveYN,
		header.TransfHUserID,
		header.TransfHLastUpdate,
		header.TransfHDataAktifYN,
		header.TransfHOrderID,
		header.TransfHSPID,
		header.TransfHPaymentMethod); err != nil {
		return errors.Wrap(err, "[DATA][InsertDataHeaderFromAPI]")
	}
	return err

}

// InsertDataDetailFromAPI ...
func (d Data) InsertDataDetailFromAPI(ctx context.Context, detail doEntity.TransfD) error {
	var err error
	if _, err = d.stmt[insertDataDetailFromAPI].ExecContext(ctx,
		detail.TransfDOutCodeTransf,
		detail.TransfDNoTransf,
		detail.TransfDGroup,
		detail.TransfDOutCodeSP,
		detail.TransfDNoSP,
		detail.TransfDProCod,
		detail.TransfDBatchNumber,
		detail.TransfDED,
		detail.TransfDQty,
		detail.TransfDQtyScan,
		detail.TransfDQtyStk,
		detail.TransfDOutCodeOrder,
		detail.TransfDNoOrder,
		detail.TransFDCategoryProduct,
		detail.TransFDEditYN,
		detail.TransfDActiveYN,
		detail.TransfDUserID,
		detail.TransfDLastUpdate,
		detail.TransfDDataAktifYN,
		detail.TransfDSPID,
		detail.TransfDSalePrice,
		detail.TransfDDiscount); err != nil {
		return errors.Wrap(err, "[DATA][InsertDataDetailFromAPI]")
	}
	return err

}

// EditDetailOrderByNoTransfandProcode ...
func (d Data) EditDetailOrderByNoTransfandProcode(ctx context.Context, noTransf string, procod string, detail doEntity.TransfD) (doEntity.TransfD, error) {
	var (
		json doEntity.TransfD
		err  error
	)
	if _, err := d.stmt[editDataDetail].ExecContext(ctx,
		detail.TransfDBatchNumber,
		detail.TransfDQtyScan,
		noTransf,
		procod); err != nil {
		return json, errors.Wrap(err, "[DATA][EditDetailOrderByNoTransfandProcode] ")
	}
	return json, err
}

// PrintReceive ...
func (d Data) PrintReceive(ctx context.Context, noTransf string, NoTranrc string) ([]pEntity.JSONPrintReceive, error) {
	var (
		rows     *sqlx.Rows
		receive  pEntity.JSONPrintReceive
		receives []pEntity.JSONPrintReceive
		err      error
	)
	rows, err = d.stmt[printReceive].QueryxContext(ctx,
		noTransf,
		NoTranrc)
	for rows.Next() {
		if err := rows.StructScan(&receive); err != nil {
			return receives, errors.Wrap(err, "[DATA][PrintReceive] ")
		}
		receives = append(receives, receive)
	}

	return receives, err
}

// GenerateReceiveHeader ...
func (d Data) GenerateReceiveHeader(ctx context.Context) (string, error) {
	var (
		newReceive string
		err        error
	)

	rows, err := d.stmt[generateReceiveHeader].QueryxContext(ctx)

	for rows.Next() {
		if err := rows.Scan(&newReceive); err != nil {
			return newReceive, errors.Wrap(err, "[DATA][GenerateReceiveHeader] ")
		}
	}
	return newReceive, err
}

// InsertReceive ...
func (d Data) InsertReceive(ctx context.Context, noReceive string, noTransf string) error {
	var (
		err error
	)
	if _, err := d.stmt[insertReceive].ExecContext(ctx,
		noReceive,
		noTransf,
	); err != nil {
		return errors.Wrap(err, "[DATA][EditDetailOrderByNoTransfandProcode] ")
	}
	return err
}

// SaveAndGenerateTranRCH ...
func (d Data) SaveAndGenerateTranRCH(ctx context.Context, header doEntity.TransfHSaveGenerate) error {
	var err error
	if _, err = d.stmt[saveAndGenerateTranRCH].ExecContext(ctx,
		header.TransfHOutCodeTransf,
		header.TransfHNoTransf,
		header.TransfHOutCodeDest,
		header.NoTranrc,
		header.TransfHTglTransf,
		header.TransfHFlag,
		header.TransfHFlagTrf,
		header.TransfHTglDwld,
		header.TransfHActiveYN,
		header.TransfHLastUpdate,
		header.TransfHDataAktifYN); err != nil {
		return errors.Wrap(err, "[DATA][SaveAndGenerateTranRCH]")
	}
	return err
}

// SaveAndGenerateTranRCD ...
func (d Data) SaveAndGenerateTranRCD(ctx context.Context, detail doEntity.TransfDSaveGenerate) error {
	var err error
	if _, err = d.stmt[saveAndGenerateTranRCD].ExecContext(ctx,
		detail.TransfDOutCodeTransf,
		detail.TransfDNoTransf,
		detail.TransfDOutCodeSP,
		detail.NoTranrc,
		detail.TransfDProCod,
		detail.TransfDBatchNumber,
		detail.TransfDQty,
		detail.TransfDQtyScan,
		detail.TransfDQtyStk,
		detail.TransfDActiveYN,
		detail.TransfDLastUpdate,
		detail.TransfDDataAktifYN); err != nil {
		return errors.Wrap(err, "[DATA][SaveAndGenerateTranRCD]")
	}
	return err
}

// GetDataDOTransfH ...
func (d Data) GetDataDOTransfH(ctx context.Context, noTransf string) (doEntity.TransfHSaveGenerate, error) {
	var (
		header doEntity.TransfHSaveGenerate
		err    error
	)

	if err := d.stmt[getDataDOTransfH].QueryRowxContext(ctx, noTransf).StructScan(&header); err != nil {
		return header, errors.Wrap(err, "[DATA][GetDataDOTransfH]")
	}

	return header, err
}

// GetDataDOTransfD ...
func (d Data) GetDataDOTransfD(ctx context.Context, noTransf string) ([]doEntity.TransfDSaveGenerate, error) {
	var (
		rows    *sqlx.Rows
		detail  doEntity.TransfDSaveGenerate
		details []doEntity.TransfDSaveGenerate
		err     error
	)
	rows, err = d.stmt[getDataDOTransfD].QueryxContext(ctx, noTransf)
	for rows.Next() {
		if err := rows.StructScan(&detail); err != nil {
			return details, errors.Wrap(err, "[DATA][GetDataDOTranfD] ")
		}
		details = append(details, detail)
	}

	return details, err
}
