package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/handlers"
)


func Health(r *gin.RouterGroup){
	handler := handlers.NewHealthHandler() 
	r.GET("/" , handler.Health)
}