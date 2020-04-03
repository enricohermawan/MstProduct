package itemPromo

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

// ItemPromo model
type ItemPromo struct {
	PromoRunningID   zero.Int    `db:"promo_runningid" json:"promo_runningid"`
	PromoCode        zero.String `db:"promo_code" json:"promo_code"`
	PromoName        zero.String `db:"promo_name" json:"promo_name"`
	PromoSpaceYN     zero.String `db:"promo_spaceyn" json:"promo_spaceyn"`
	PromoAllOutletYN zero.String `db:"promo_alloutletyn" json:"promo_alloutletyn"`
	LastUpdate       time.Time   `db:"lastupdate" json:"lastupdate"`
	Host             zero.String `db:"host" json:"host"`
	Pph              zero.Float  `db:"pph" json:"pph"`
	PromoEdlpYN      zero.String `db:"promo_edlpyn" json:"promo_edlpyn"`
	PromoTipeBayar   zero.Int    `db:"promo_tipebayar" json:"promo_tipebayar"`
	PromoLeadTime    zero.Int    `db:"promo_leadtime" json:"promo_leadtime"`
	PromoDept        zero.String `db:"promo_dept" json:"promo_dept"`
	PromoAlokasiYN   zero.String `db:"promo_alokasiyn" json:"promo_alokasiyn"`
	PromoActiveYN    zero.String `db:"promo_activeyn" json:"promo_activeyn"`
}
