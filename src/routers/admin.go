package routers

import (
	"snapp/handlers"
	"snapp/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminsRouter(r *gin.RouterGroup) {
	h := handlers.GetAdminHelper()
	r.GET("/see/users", middlewares.CheckAdmin, h.SeeUsers)
}