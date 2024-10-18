package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()
	r.GET("/", h.Test)
	r.POST("/header1", h.HeaderBinder1)
	r.POST("/header2", h.HeaderBinder2)
	r.POST("/query1", h.QueryBinder1)
	r.POST("/query2", h.QueryBinder2)
	r.POST("/params/:id/:name/:family", h.ParamsBinder1)
	r.POST("/body", h.BodyBinder)
	r.POST("/form", h.FormBinder)
	r.POST("/file", h.FileBinder)

}
