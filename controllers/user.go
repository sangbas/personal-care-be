package controllers

import (
	"net/http"
	"personal-care-be/forms"
	"personal-care-be/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

//One ...
func (ctrl UserController) One(c *gin.Context) {
	// userID := getUserID(c)
	id := c.Param("userId")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := userModel.One(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "User not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
	}
}

//Create ...
func (ctrl UserController) Create(c *gin.Context) {
	var user forms.User
	c.ShouldBind(&user)

	id, err := userModel.Create(user)

	if err != nil {
		panic(err.Error())
	} else {
		c.JSON(201, gin.H{
			"id": id,
		})
	}

}
