package routers

import (
	"snapp/handlers"

	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.RouterGroup) {
	h := handlers.GetUserHelper()
	r.POST("/new/user", h.AddUser)
}