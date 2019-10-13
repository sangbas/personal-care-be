package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//DB ...
type DB struct {
	*sql.DB
}

//config ...
const (
	DbUser     = "b980eb19293673"
	DbPassword = "36d1a594"
	DbName     = "heroku_9532b652abd0830"
	DbHost     = "us-cdbr-iron-east-02.cleardb.net"
)

// var db *gorp.DbMap

//Init ...
// func Init() {

// 	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
// 	// 	DbUser, DbPassword, DbName)

// 	db, err := sql.Open("mysql", "b980eb19293673:36d1a594@tcp(us-cdbr-iron-east-02.cleardb.net)/heroku_9532b652abd0830")
// 	// if there is an error opening the connection, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// }

//CreateCon Create MySQL connection ...
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", DbUser+":"+DbPassword+"@tcp("+DbHost+")/"+DbName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

//ConnectDB ...
// func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
// 	db, err := sql.Open("postgres", dataSourceName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	defer db.close()
// 	return dbmap, nil
// }

// //GetDB ...
// func GetDB() *gorp.DbMap {
// 	return db
// }
