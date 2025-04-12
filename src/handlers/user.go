package handlers

import (
	"fmt"
	"net/http"
	"snapp/data"
	"snapp/responses"

	"github.com/gin-gonic/gin"
)

type UserHelper struct {
}

func GetUserHelper() *UserHelper {
	return &UserHelper{}
}

func (h UserHelper) AddUser(c *gin.Context) {
	user := data.SimpleUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	ok := data.AddUser(user)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			fmt.Errorf("this ID has used"),
		))
		return
	}
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"new user added",
	))
}
