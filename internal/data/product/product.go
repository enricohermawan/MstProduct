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
