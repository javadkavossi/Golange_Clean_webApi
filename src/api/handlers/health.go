package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))

}

func (h *HealthHandler) HealthPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, " ")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Working! Post", true, 0))

}
func (h *HealthHandler) HealthPostById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(fmt.Sprintf("Working! id : %s", id), true, 0))

}
