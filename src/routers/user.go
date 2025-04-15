package routers

import (
	"snapp/handlers"

	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.RouterGroup) {
	h := handlers.GetUserHelper()
	r.POST("/new/user", h.NewUser)
	r.POST("/delete/account", h.DeleteUserAccount)
	r.POST("/charge/account", h.ChargeAccount)
	r.GET("/view", h.View)
	r.POST("/order", h.Order)
}