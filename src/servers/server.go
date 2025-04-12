package servers

import (
	"fmt"
	"snapp/config"
	"snapp/routers"
	"snapp/validations"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewServer(cfg *config.Config) {
	engine := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("id", validations.IdValidator)
		val.RegisterValidation("password", validations.PasswoordValidator)
	}

	engine.Use(gin.Recovery(), gin.Logger())

	snapp := engine.Group("/snapp")
	{
		// driver := snapp.Group("/driver")
		user := snapp.Group("/user")
		routers.UsersRouter(user)
		admin := snapp.Group("/admin")
		routers.AdminsRouter(admin)
	}

	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}