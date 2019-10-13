package controllers

import (
	"net/http"
	"personal-care-be/forms"
	"personal-care-be/models"

	"github.com/gin-gonic/gin"
)

//OrderController ...
type OrderController struct{}

var orderModel = new(models.OrderModel)

//GetByCategory ...
func (ctrl OrderController) GetByCategory(c *gin.Context) {
	category := c.Param("category")
	data, err := orderModel.ListByCategory(category)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
}

//Create ...
func (ctrl OrderController) Create(c *gin.Context) {
	var order forms.Order
	c.ShouldBind(&order)

	id, err := orderModel.Create(order)

	if err != nil {
		panic(err.Error())
	} else {
		c.JSON(201, gin.H{
			"id": id,
		})
	}
}
