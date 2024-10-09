package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/middlewares"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/routers"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/validations"
	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobil", validations.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validations.PasswordValidator, true)

	}
	r.Use(middlewares.Cors(cfg))

	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /*middlewares.TestMiddleware()*/, middlewares.LimiterByRequest())

	api := r.Group("/api")

	v1 := api.Group("/v1")

	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))




}
