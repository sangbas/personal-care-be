package models

import (
	"personal-care-be/db"
	"personal-care-be/forms"
)

//User ...
type User struct {
	UserID int64  `db:"user_id, primarykey, autoincrement" json:"user_id"`
	Name   string `db:"name" json:"name"`
	Photo  string `db:"photo" json:"photo"`
}

//UserModel ...
type UserModel struct{}

// var con *sql.DB

//Create User
func (m UserModel) Create(form forms.User) (userID int64, err error) {
	con := db.CreateCon()

	res, err := con.Exec(`INSERT INTO user (name,photo) VALUES ( ?, '' ) `, form.Name)
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

//One Get User By ID ...
func (m UserModel) One(id int64) (user User, err error) {
	con := db.CreateCon()
	sqlStatement := `SELECT * FROM user where user_id = ?;`
	row := con.QueryRow(sqlStatement, id)
	err = row.Scan(&user.UserID, &user.Name, &user.Photo)
	// err = con.SelectOne(&user, "SELECT name, photo FROM user where user_id = ?", c.Param("userId"))
	return user, err
}
