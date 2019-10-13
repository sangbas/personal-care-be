package main

import (
	"personal-care-be/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	user := new(controllers.UserController)
	r.GET("/user/:userId", user.One)
	r.POST("/register", user.Create)

	order := new(controllers.OrderController)
	r.GET("/order/:category", order.GetByCategory)
	r.POST("/order", order.Create)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  "404",
			"message": "Not Found.",
		})
	})

	r.Run()
}
