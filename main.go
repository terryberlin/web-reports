package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	//"github.com/terryberlin/web-reports/reportstructs/operationsreports"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func main() {
	http.HandleFunc("/operations", Operations)
	http.HandleFunc("/coupons", Coupons)
	http.HandleFunc("/voids", Voids)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Println(err)
	}
}

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

//Void is a structure for void data.
type (
	Void struct {
		UnitID            *string `db:"unit_id" json:"unitID"`
		Store             *string `db:"store" json:"store"`
		StoreName         *string `db:"store_name" json:"storeName"`
		Date              *string `db:"date" json:"date"`
		Sales             *string `db:"sales" json:"sales"`
		Guests            *string `db:"guests" json:"guests"`
		CancelItems       *string `db:"cancel_items" json:"cancelItems"`
		CancelItemsCount  *string `db:"cancel_items_count" json:"cancelItemsCount"`
		CancelOrders      *string `db:"cancel_orders" json:"cancelOrders"`
		CancelOrdersCount *string `db:"cancel_orders_count" json:"cancelOrdersCount"`
		Deletes           *string `db:"deletes" json:"deletes"`
		DeletesCount      *string `db:"deletes_count" json:"deletesCount"`
		Voids             *string `db:"voids" json:"voids"`
		VoidsCount        *string `db:"voids_count" json:"voidsCount"`
		Refunds           *string `db:"refunds" json:"refunds"`
		RefundsCount      *string `db:"refunds_count" json:"refundsCount"`
	}
)

//Operations is a function for writing and requesting date data.  The url accepts arguments for begin date, end date, and store units.
func Operations(w http.ResponseWriter, r *http.Request) {
	begdate := r.URL.Query().Get("begdate")
	enddate := r.URL.Query().Get("enddate")
	units := r.URL.Query().Get("units")
	sql := `Exec spQsOperationsReport $1, $2, Null, Null, Null, Null, Null, Null, Null, Null, 'units', $3`
	operations := []Operation{}
	err := DB().Select(&operations, sql, begdate, enddate, units)
	if err != nil {
		log.Println(err)
	}
	json, err := json.Marshal(operations)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(json))
}

//Coupons is a function for writing and requestion coupons data.  The url accepts arguments for begin date, end date, and store units.
func Coupons(w http.ResponseWriter, r *http.Request) {
	begdate := r.URL.Query().Get("begdate")
	enddate := r.URL.Query().Get("enddate")
	units := r.URL.Query().Get("units")

	sql := `Exec spQsDiscountsReport $1, $2, Null, Null, Null, Null, Null, Null, Null, Null, 'units', $3`
	coupons := []Coupon{}
	err := DB().Select(&coupons, sql, begdate, enddate, units)
	if err != nil {
		log.Println(err)
	}
	json, err := json.Marshal(coupons)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(json))
}

//Voids is a function for writing and requestion void data.  The url accepts arguments for begin date, end date, and store units.
func Voids(w http.ResponseWriter, r *http.Request) {
	begdate := r.URL.Query().Get("begdate")
	enddate := r.URL.Query().Get("enddate")
	units := r.URL.Query().Get("units")

	sql := `Exec getVoids $1, $2, Null, Null, Null, Null, Null, Null, Null, Null, 'units', $3`
	voids := []Void{}
	err := DB().Select(&voids, sql, begdate, enddate, units)
	if err != nil {
		log.Println(err)
	}
	json, err := json.Marshal(voids)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(json))
}

//DB is a function that connects to SQL server.
func DB() *sqlx.DB {
	pass := os.Getenv("DB_PASS")
	db, err := sqlx.Connect("mssql", fmt.Sprintf("server=192.168.1.34;user id=sa;password=%s;database=quikserve;log64;encrypt=disable", pass))
	if err != nil {
		log.Println(err)
	}
	return db
}

// func DB() *sqlx.DB {
// 	db, err := sqlx.Connect("mssql", "server=192.168.1.34;user id=sa;password=1234;database=quikserve;log64;encrypt=disable")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return db
// }

/*
http://localhost:5000/operations?begdate=10/01/2017&enddate=10/02/2017&units=2,3,5,6

http://localhost:5000/coupons?begdate=10/01/2017&enddate=10/02/2017&units=2,3,5,6

http://localhost:5000/voids?begdate=10/01/2017&enddate=10/02/2017&units=2,3,5,6
*/
