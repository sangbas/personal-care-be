package models

import (
	"personal-care-be/db"
	"personal-care-be/forms"
)

//Order ...
type Order struct {
	OrderID  int64  `db:"order_id, primarykey, autoincrement" json:"order_id"`
	UserID   int64  `db:"user_id" json:"-"`
	Beverage string `db:"beverage" json:"beverage"`
	Served   int    `db:"served" json:"served"`
}

//OrderModel ...
type OrderModel struct{}

// var con *sql.DB

//Create Model ...
func (m OrderModel) Create(form forms.Order) (OrderID int64, err error) {
	con := db.CreateCon()

	res, err := con.Exec("INSERT INTO `order` (user_id,beverage) VALUES ( ?, ? ) ", form.UserID, form.Beverage)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	} else {
		id, err := res.LastInsertId()

		if err != nil {
			println("Error:", err.Error())
		} else {
			println("LastInsertId:", id)
			return id, err
		}
	}
	return 0, err
}

//ListByCategory ...
func (m OrderModel) ListByCategory(category string) (orders []forms.OrderResp, err error) {
	con := db.CreateCon()

	sqlStatement := ""
	if category == "beverage" {
		sqlStatement = "SELECT o.order_id, u.name, o.beverage FROM `order` o JOIN `user` u ON u.user_id=o.user_id where beverage != 'undefined' order by order_id asc;"
	} else {
		sqlStatement = "SELECT o.order_id, u.name, o.beverage FROM `order` o JOIN `user` u ON u.user_id=o.user_id where beverage = 'undefined' order by order_id asc;"
	}

	row, err := con.Query(sqlStatement)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for row.Next() {
		var order forms.OrderResp
		row.Scan(&order.OrderID, &order.User, &order.Beverage)
		orders = append(orders, order)
	}
	return orders, err
}
