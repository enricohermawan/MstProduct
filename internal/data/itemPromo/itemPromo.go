package itemPromo

import (
	"context"
	itemPromoEntity "go-itempromo/internal/entity/itemPromo"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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

// Tambahkan query di dalam const
// getAllUser = "GetAllUser"
// qGetAllUser = "SELECT * FROM users"
const (
	getItemPromo  = "GetItemPromo"
	qGetItemPromo = "SELECT * FROM m_itempromo"
)

// Tambahkan query ke dalam key value order agar menjadi prepared statements
// readStmt = []statement{
// 	{getAllUser, qGetAllUser},
// }
var (
	readStmt = []statement{
		{getItemPromo, qGetItemPromo},
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

// GetItemPromo ..
func (d Data) GetItemPromo(ctx context.Context) ([]itemPromoEntity.ItemPromo, error) {
	var (
		promoItem  itemPromoEntity.ItemPromo
		promoItems []itemPromoEntity.ItemPromo
		err        error
	)

	// Query ke database
	rows, err := d.stmt[getItemPromo].QueryxContext(ctx)

	// Looping seluruh row data
	for rows.Next() {
		// Insert row data ke struct user
		if err := rows.StructScan(&promoItem); err != nil {
			return promoItems, errors.Wrap(err, "[DATA][GetItemPromo] ")
		}
		// Tambahkan struct user ke array user
		promoItems = append(promoItems, promoItem)
	}
	// Return users array
	return promoItems, err
}
