package couponreport

//Coupon is a structure for coupon data.
type (
	Coupon struct {
		UnitID    *string `db:"unit_id" json:"unitID"`
		Store     *string `db:"store" json:"store"`
		StoreName *string `db:"store_name" json:"storeName"`
		Date      *string `db:"date" json:"date"`
		Descr     *string `db:"descr" json:"descr"`
		Group     *string `db:"group" json:"group"`
		Amount    *string `db:"amount" json:"amount"`
		Quantity  *string `db:"quantity" json:"quantity"`
		Sales     *string `db:"sales" json:"sales"`
		Percent   *string `db:"percent" json:"percent"`
	}
)
