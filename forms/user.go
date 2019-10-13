package forms

//User ...
type User struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Photo string `form:"photo" json:"photo" binding:"required"`
}
