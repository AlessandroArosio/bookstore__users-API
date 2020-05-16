package users

import (
	"github.com/alessandroarosio/bookstore_users-API/domain/users"
	"github.com/alessandroarosio/bookstore_users-API/services"
	"github.com/alessandroarosio/bookstore_users-API/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
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

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "todo")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "todo")
}
