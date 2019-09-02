package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"


// Binding from JSON
type User struct {
	UserId 	string `form:"userId" json:"userId"  binding:"required"`
	Name 	string `form:"name" json:"name" binding:"required"`
	Photo 	string `form:"photo" json:"photo" binding:"required"`
	AreaId  []string `json:"checkin" binding:"required"`
}

type CheckIn struct {
	UserId 	string `form:"userId" json:"userId"  binding:"required"`
	AreaId 	string `form:"areaId" json:"areaId" binding:"required"`
}

type Order struct {
	UserId 	string `form:"userId" json:"userId"  binding:"required"`
	BeverageId 	string `form:"beverageId" json:"beverageId" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	db, err := sql.Open("mysql", "b980eb19293673:36d1a594@tcp(us-cdbr-iron-east-02.cleardb.net)/heroku_9532b652abd0830")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	r.GET("/user/:userId", func(c *gin.Context) {
		var user User
		// Execute the query
		err = db.QueryRow("SELECT name, photo FROM user where user_id = ?", c.Param("userId")).Scan(&user.Name, &user.Photo)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.JSON(200, gin.H{
			"userId": c.Param("userId"),
			"name": user.Name,
			"photo": user.Photo,
		})
	})
	r.POST("/register", func(c *gin.Context) {
		var user User
		c.ShouldBind(&user)

		// perform a db.Query insert
		res, err := db.Exec(`INSERT INTO user (name,photo) VALUES ( ?, '' ) `, user.Name)
		// defer res.Close()
		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		} else {
			id, err := res.LastInsertId()
			
			if err != nil {
				println("Error:", err.Error())
			} else {
				println("LastInsertId:", id)
				c.JSON(200, gin.H{
					"id": id,
				})
			}
		}
		// be careful deferring Queries if you are using transactions
		// defer res.Close()

		
	})
	
	r.POST("/checkin", func(c *gin.Context) {
		var checkIn CheckIn
		c.ShouldBind(&checkIn)
		// perform a db.Query insert
		insert, err := db.Query("INSERT INTO check_in (user_id, area_id) VALUES ( "+checkIn.UserId+", "+checkIn.AreaId+" )")

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
		c.JSON(200, gin.H{
			"message": checkIn.UserId,
		})
	})

	r.POST("/order", func(c *gin.Context) {
		var order Order
		c.ShouldBind(&order)
		// perform a db.Query insert
		insert, err := db.Query("INSERT INTO `order` (user_id, beverage_id, served) VALUES ( "+order.UserId+", "+order.BeverageId+", 0 )")

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
		c.JSON(200, gin.H{
			"message": order.UserId,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

