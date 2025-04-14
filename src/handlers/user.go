package handlers

import (
	"fmt"
	"net/http"
	"snapp/config"
	"snapp/data"
	"snapp/loggers"
	"snapp/responses"

	"github.com/gin-gonic/gin"
)

var log = loggers.NewLogger(config.GetConfig())
type UserHelper struct {
}

func GetUserHelper() *UserHelper {
	return &UserHelper{}
}

type SimpleUser struct {
	Fullname string `json:"fullname" binding:"required,min=5,max=25"`
	ID       string `json:"id" binding:"required,id"`
	Password string `json:"password" binding:"required,password"`
}

func AddUser(user SimpleUser) bool {
	for k := range data.Users {
		if k == user.ID {
			return false
		}
	}
	new := data.User{
		Fullname: user.Fullname,
		Id:       user.ID,
		Password: user.Password,
	}
	new.Wallet = 0
	data.Users[new.Id] = new
	return true
}

// @Summary New User
// @Description New User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /user/new/user [post]
func (h UserHelper) AddUser(c *gin.Context) {
	log := loggers.NewLogger(config.GetConfig())
	user := SimpleUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	ok := AddUser(user)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			fmt.Errorf("this ID has used"),
		))
		return
	}
	log.Info(loggers.User, loggers.NewUser, "new user created", map[loggers.ExtraKey]interface{}{loggers.UserID: user.ID, loggers.Fullname: user.Fullname})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"new user added",
	))
}

type DeleteSampleUser struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func DeleteUser(target DeleteSampleUser) error {
	for k, v := range data.Users {
		if k == target.Id {
			if v.Password == target.Password {
				log.Info(loggers.User, loggers.Delete, "user account delete", nil)
				delete(data.Users, k)
				return nil
			}
		}
	}
	return fmt.Errorf("there is no user with this id and password")
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /user/delete/account [post]
func (h *UserHelper) DeleteAccount(c *gin.Context) {
	target_user := DeleteSampleUser{}
	err := c.ShouldBindJSON(&target_user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	err = DeleteUser(target_user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"user account deleted",
	))
}
