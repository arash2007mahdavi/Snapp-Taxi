package handlers

import (
	"fmt"
	"net/http"
	"snapp/data"
	"snapp/loggers"
	"snapp/responses"

	"github.com/gin-gonic/gin"
)

type AdminHelper struct {
}

func GetAdminHelper() *AdminHelper {
	return &AdminHelper{}
}

// @Summary See Drivers
// @Description See Drivers
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /admin/see/drivers [get]
func (h AdminHelper) SeeDrivers(c *gin.Context) {
	log.Info(loggers.Admin, loggers.SeeDrivers, "admin see drivers", map[loggers.ExtraKey]interface{}{loggers.AdminID:c.GetHeader("id")})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		data.Drivers,
	))	
}

// @Summary See Users
// @Description See Users
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /admin/see/users [get]
func (h AdminHelper) SeeUsers(c *gin.Context) {
	log.Info(loggers.Admin, loggers.SeeUsers, "admin see users", map[loggers.ExtraKey]interface{}{loggers.AdminID:c.GetHeader("id")})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		data.Users,
	))	
}

// @Summary See Admins
// @Description See Admins
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /admin/see/admins [get]
func (h AdminHelper) SeeAdmins(c *gin.Context) {
	log.Info(loggers.Admin, loggers.SeeAdmins, "admin see admins", map[loggers.ExtraKey]interface{}{loggers.AdminID:c.GetHeader("id")})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		data.Admins,
	))	
}

type NewAdmin struct {
	Fullname string `json:"fullname" binding:"required,min=5,max=25"`
	Id       string `json:"id" binding:"required,id"`
	Password string `json:"password" binding:"required,password"`
}

func AddAdmin(new NewAdmin) error {
	for k := range data.Admins {
		if k == new.Id {
			log.Error(loggers.Admin, loggers.Add, "invalid ID", nil)
			return fmt.Errorf("there is an admin with this ID")
		}
	}
	newadmin := data.Admin{
		Fullname: new.Fullname,
		Id: new.Id,
		Password: new.Password,
	}
	data.Admins[newadmin.Id] = newadmin
	return nil
}

// @Summary New Admin
// @Description New Admin
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /admin/new/admin [post]
func (h AdminHelper) NewAdmin(c *gin.Context) {
	new := NewAdmin{}
	err := c.ShouldBindJSON(&new)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	err = AddAdmin(new)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(
			false,
			http.StatusNotAcceptable,
			err,
		))
		return
	}
	log.Info(loggers.Admin, loggers.NewAdmin, "new admin created", map[loggers.ExtraKey]interface{}{loggers.AdminID: new.Id, loggers.Fullname: new.Fullname})
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(
		true,
		http.StatusOK,
		"new admin added",
	))
}