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
	tampilDetailDataByNoReceive  = "TampilDetailDataByNoReceive"
	qtampilDetailDataByNoReceive = "SELECT * FROM MH_Product_Alvin JOIN TranRCD ON MH_Product_Alvin.Pro_Code = TranRCD.TranrcD_ProCod WHERE TranRCD.TranrcD_NoTranrc = ?"
)

var (
	readStmt = []statement{
		{tampilDetailDataByNoReceive, qtampilDetailDataByNoReceive},
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

// TampilDetailReceiveByNoReceive ..
func (d Data) TampilDetailReceiveByNoReceive(ctx context.Context, NoTranrc string) (pEntity.MstProduct, error) {
	var (
		product pEntity.Gabung
		err     error
	)

	// Query ke database
	err = d.stmt[tampilDetailDataByNoReceive].QueryRowxContext(ctx, NoTranrc).StructScan(&product)

	if err != nil {
		return product.MstProduct, errors.Wrap(err, "[DATA][TampilDetailReceiveByNoReceive]")
	}

	return product.MstProduct, err
}
