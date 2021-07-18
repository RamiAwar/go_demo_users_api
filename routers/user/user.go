package user

import (
	"net/http"

	"github.com/RamiAwar/go_demo_users_api/models/user"
	"github.com/RamiAwar/go_demo_users_api/services"
	"github.com/RamiAwar/go_demo_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError(err)
		c.JSON(restError.Status, restError)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
}

func AutocompleteUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
}