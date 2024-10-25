package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/javadkavossi/Golange_Clean_webApi/src/api/middlewares"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/routers"
	validation "github.com/javadkavossi/Golange_Clean_webApi/src/api/validations"

	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/docs"
	"github.com/javadkavossi/Golange_Clean_webApi/src/pkg/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	// cfg := config.GetConfig()
	r := gin.New()
	// *-------------------------- Get Functions
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	// ?-------------------------------------------------

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /*middlewares.TestMiddleware()*/, middlewares.LimiterByRequest())

	RegisterValidators()
	RegisterSwagger(r, cfg)
	RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

// *-------------------------- Functions

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	api := r.Group("/api")

	v1 := api.Group("/v1")

	{
		health := v1.Group("/health")
		users := v1.Group("/users")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
		// User
		routers.User(users, cfg)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

// ?-------------------------------------------------

// * ------------------- Swagger

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
