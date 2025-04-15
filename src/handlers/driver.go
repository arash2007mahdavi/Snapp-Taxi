package handlers

import (
	"fmt"
	"net/http"
	"snapp/data"
	"snapp/loggers"
	"snapp/responses"

	"github.com/gin-gonic/gin"
)

type DriverHelper struct {
}

func GetDriverHelper() *DriverHelper {
	return &DriverHelper{}
}

type SampleDriver struct {
	Fullname string `json:"fullname" binding:"required,min=5,max=25"`
	Id       string `json:"id" binding:"required,id"`
	Password string `json:"password" binding:"required,password"`
}

func AddDriverToDataBase(driver SampleDriver) bool {
	for k := range data.Drivers {
		if driver.Id == k {
			return false
		}
	}
	new := data.Driver{
		Fullname: driver.Fullname,
		Id: driver.Id,
		Password: driver.Password,
		Wallet: 0,
	}
	data.AddToDataBase(new)
	return true
}

// @Summary New Driver
// @Description New Driver
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /driver/new/driver [post]
func (h *DriverHelper) NewDriver(c *gin.Context) {
	new := &SampleDriver{}
	err := c.ShouldBindJSON(&new)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	ok := AddDriverToDataBase(*new)
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			fmt.Errorf("this ID has used"),
		))
		return
	}
	log.Info(loggers.Driver, loggers.NewDriver, "new driver added", map[loggers.ExtraKey]interface{}{loggers.AdminID: new.Id, loggers.Fullname: new.Fullname})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"new driver added",
	))
}

func DeleteDriverFromDataBase(target DeleteSampleUser) error {
	for k, v := range data.Drivers {
		if k == target.Id {
			if v.Password == target.Password {
				log.Info(loggers.Driver, loggers.Delete, "driver account delete", nil)
				delete(data.Drivers, k)
				return nil
			}
		}
	}
	return fmt.Errorf("there is no driver with this id and password")
}

// @Summary Delete Driver
// @Description Delete Driver
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /driver/delete/account [post]
func (h *DriverHelper) DeleteDriverAccount(c *gin.Context) {
	target_driver := DeleteSampleUser{}
	err := c.ShouldBindJSON(&target_driver)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	err = DeleteDriverFromDataBase(target_driver)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	log.Info(loggers.Driver, loggers.Delete, "driver account deleted", map[loggers.ExtraKey]interface{}{loggers.DriverID: target_driver.Id})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"driver account deleted",
	))
}

// @Summary See Travels Line
// @Description See Travels Line
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /driver/see/travels/line [get]
func (h *DriverHelper) SeeTravelsLine(c *gin.Context) {
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		data.LineTravels,
	))
}