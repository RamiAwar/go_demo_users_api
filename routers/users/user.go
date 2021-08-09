package users

import (
	"net/http"
	"strconv"

	"github.com/RamiAwar/go_demo_users_api/models/users"
	s "github.com/RamiAwar/go_demo_users_api/services"
	"github.com/RamiAwar/go_demo_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, err := s.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshal())
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		userErr := errors.BadRequestError("Invalid user id")
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := s.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshal())
}

func SearchUser(c *gin.Context) {
	users, err := s.UserService.Search(c.Query("q"))
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, users.Marshal())
}
