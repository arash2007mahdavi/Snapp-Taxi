package routers

import (
	"snapp/handlers"
	"snapp/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminsRouter(r *gin.RouterGroup) {
	h := handlers.GetAdminHelper()
	r.POST("/new/admin", middlewares.CheckAdmin, h.NewAdmin)
	r.GET("/see/users", middlewares.CheckAdmin, h.SeeUsers)
	r.GET("/see/admins", middlewares.CheckAdmin, h.SeeAdmins)
}