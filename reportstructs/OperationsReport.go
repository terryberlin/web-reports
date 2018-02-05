package operationsreport

//Operation is a structure for operations data.
type (
	Operation struct {
		UnitID                  *string `db:"unit_id" json:"unitID"`
		Store                   *string `db:"store" json:"store"`
		StoreName               *string `db:"store_name" json:"storeName"`
		Date                    *string `db:"date" json:"date"`
		AdjustedSales           *string `db:"adjusted_sales" json:"adjustedSales"`
		AdjustedSalesLYDate     *string `db:"adjusted_sales_ly_date" json:"adjustedSalesLYDate"`
		AdjustedSalesCYDateComp *string `db:"adjusted_sales_cy_date_comp" json:"adjustedSalesCYDateComp"`
		AdjustedSalesLYDateComp *string `db:"adjusted_sales_ly_date_comp" json:"adjustedSalesLYDateComp"`
		AdjustedSalesLYDOW      *string `db:"adjusted_sales_ly_dow" json:"adjustedSalesLYDOW"`
		AdjustedSalesCYDOWComp  *string `db:"adjusted_sales_cy_dow_comp" json:"adjustedSalesCYDOWComp"`
		AdjustedSalesLYDOWComp  *string `db:"adjusted_sales_ly_dow_comp" json:"adjustedSalesLYDOWComp"`
	}
)
