package routers

import (
	"snapp/handlers"

	"github.com/gin-gonic/gin"
)

func DriverRouter(r *gin.RouterGroup) {
	h := handlers.GetDriverHelper()
	r.POST("/new/driver", h.NewDriver)
	r.POST("/delete/account", h.DeleteDriverAccount)
	r.GET("/see/travels/line", h.SeeTravelsLine)
}