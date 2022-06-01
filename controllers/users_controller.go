package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JAlayon/bookstore_users-api/domain/users"
	"github.com/JAlayon/bookstore_users-api/errors"
	"github.com/JAlayon/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userError != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	result, getUserError := services.GetUser(userId)
	if getUserError != nil {
		c.JSON(getUserError.Status, getUserError)
		return
	}
	c.JSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
