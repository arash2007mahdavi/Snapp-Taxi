package handlers

import "github.com/gin-gonic/gin"

type AdminHelper struct {
}

func GetAdminHelper() *AdminHelper {
	return &AdminHelper{}
}

func (h AdminHelper) SeeUsers(c *gin.Context) {
	
}