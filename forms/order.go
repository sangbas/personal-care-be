package forms

//Order ...
type Order struct {
	UserID   string `form:"userId" json:"userId"  binding:"required"`
	Beverage string `form:"beverage" json:"beverage" binding:"required"`
}

//OrderResp ...
type OrderResp struct {
	OrderID  int64  `form:"orderId" json:"orderId"  binding:"required"`
	User     string `form:"user" json:"user"  binding:"required"`
	Beverage string `form:"beverage" json:"beverage" binding:"required"`
}
