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

func AddUserToDataBase(user SimpleUser) bool {
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
func (h UserHelper) NewUser(c *gin.Context) {
	new := &SimpleUser{}
	err := c.ShouldBindJSON(&new)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	ok := AddUserToDataBase(*new)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			fmt.Errorf("this ID has used"),
		))
		return
	}
	log.Info(loggers.User, loggers.NewUser, "new user added", map[loggers.ExtraKey]interface{}{loggers.UserID: new.ID, loggers.Fullname: new.Fullname})
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

func DeleteUserFromDataBase(target DeleteSampleUser) error {
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
func (h *UserHelper) DeleteUserAccount(c *gin.Context) {
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
	err = DeleteUserFromDataBase(target_user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	log.Info(loggers.Driver, loggers.Delete, "user account deleted", map[loggers.ExtraKey]interface{}{loggers.UserID: target_user.Id})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"user account deleted",
	))
}

type SampleOrder struct {
	CustomerID string    `json:"customer" binding:"required"`
	Genesis    data.City `json:"genesis" binding:"required"`
	Target     data.City `json:"target" binding:"required"`
}

func NewLineTravel(order SampleOrder, distance float64, price float64) data.LineTravel {
	lineTravel := &data.LineTravel{
		CustomerID: order.CustomerID,
		Genesis:    order.Genesis,
		Target:     order.Target,
		Distance:   distance,
		Price:      price,
	}
	data.AddToLineTravels(*lineTravel)
	return *lineTravel
}

// @Summary Order
// @Description Order
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /user/order [post]
func (h *UserHelper) Order(c *gin.Context) {
	order := &SampleOrder{}
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	ok1 := data.ValidateCity(string(order.Genesis))
	ok2 := data.ValidateCity(string(order.Target))
	if !ok1 || !ok2 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			fmt.Errorf("invalid citys"),
		))
		return
	}
	distance := data.GetDistance(order.Genesis, order.Target)
	price := data.GetPrice(distance)
	for k, v := range data.Users {
		if order.CustomerID == k {
			if price <= v.Wallet {
				lineTravel := NewLineTravel(*order, distance, price)
				c.JSON(http.StatusOK, responses.GenerateNormalResponse(
					true,
					http.StatusOK,
					lineTravel,
				))
				return
			}
			c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
				false,
				http.StatusNotAcceptable,
				fmt.Errorf("your money isnt enoght for order"),
			))
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
		false,
		http.StatusNotAcceptable,
		fmt.Errorf("couldnt find any user with this ID"),
	))
}

type SampleChargeAccount struct {
	Id     string  `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

// @Summary Charge Account
// @Description Charge Account
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /user/charge/account [post]
func (h *UserHelper) ChargeAccount(c *gin.Context) {
	sample := &SampleChargeAccount{}
	err := c.ShouldBindJSON(&sample)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	for k := range data.Users {
		if k == sample.Id {
			user := data.Users[k]
			user.Wallet += sample.Amount
			data.Users[k] = user
			c.JSON(http.StatusOK, responses.GenerateNormalResponse(
				true,
				http.StatusOK,
				"charged successfuly",
			))
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
		false,
		http.StatusNotAcceptable,
		fmt.Errorf("there is no user with this ID"),
	))
}

type SampleView struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary View
// @Description View
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /user/view [get]
func (h *UserHelper) View(c *gin.Context) {
	sample := &SampleView{}
	err := c.ShouldBindJSON(&sample)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false, 
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	for k, v := range data.Users {
		if k == sample.Id {
			if v.Password == sample.Password {
				c.JSON(http.StatusOK, responses.GenerateNormalResponse(
					true,
					http.StatusOK,
					v,
				))
				return
			}
			c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
				false,
				http.StatusNotAcceptable,
				fmt.Errorf("invalid password"),
			))
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
		false,
		http.StatusNotAcceptable,
		fmt.Errorf("there is no user with this ID"),
	))
}
