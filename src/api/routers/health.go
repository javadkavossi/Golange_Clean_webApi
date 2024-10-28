package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/handler"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()
	r.GET("/", handler.Health)

}
