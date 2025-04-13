package middlewares

import (
	"fmt"
	"net/http"
	"snapp/config"
	"snapp/data"
	"snapp/loggers"
	"snapp/responses"

	"github.com/gin-gonic/gin"
)

var logger = loggers.NewLogger(config.GetConfig())

func CheckAdmin(c *gin.Context) {
	id := c.GetHeader("id")
	password := c.GetHeader("password")
	ok := data.AdminCheck(id, password)
	if !ok {
		logger.Warn(loggers.Admin, loggers.Check, "invalid Admin", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.GenerateResponseWithError(
			false,
			http.StatusUnauthorized,
			fmt.Errorf("invalid Admin"),
		))
		return
	}
	c.Next()
}