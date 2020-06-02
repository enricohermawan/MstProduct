package product

import (
	"context"
	"log"
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
)

var (
	readStmt = []statement{

		{getAllDetailReceive, qgetAllDetailReceive},
		{getAllHeaderReceive, qgetAllHeaderReceive},
		{getDataDetailByNoReceive, qgetDataDetailByNoReceive},
		{getDataHeaderByNoReceive, qgetDataHeaderByNoReceive},
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
