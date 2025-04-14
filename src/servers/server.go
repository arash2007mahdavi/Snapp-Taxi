package servers

import (
	"fmt"
	"snapp/docs"
	"snapp/config"
	"snapp/routers"
	"snapp/validations"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		driver := snapp.Group("/driver")
		routers.DriverRouter(driver)
		user := snapp.Group("/user")
		routers.UsersRouter(user)
		admin := snapp.Group("/admin")
		routers.AdminsRouter(admin)
	}

	docs.SwaggerInfo.Title = "Snapp"
	docs.SwaggerInfo.Description = "Snapp Taxi"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/snapp"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", cfg.Server.Port)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}