package router

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/handler"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/middlewares"
	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
