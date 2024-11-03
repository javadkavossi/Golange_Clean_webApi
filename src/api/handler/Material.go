package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/dto"
	_ "github.com/javadkavossi/Golange_Clean_webApi/src/api/helper"
	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/dependency"
	_ "github.com/javadkavossi/Golange_Clean_webApi/src/domain/filter"
	"github.com/javadkavossi/Golange_Clean_webApi/src/usecase"
)

type MaterialHandler struct {
	usecase *usecase.MaterialUsecase
}

func NewMaterialHandler(cfg *config.Config) *MaterialHandler {
	return &MaterialHandler{
		usecase: usecase.NewMaterialUsecase(cfg, dependency.GetMaterialRepository(cfg)),
	}
}

// CreateColor godoc
// @Summary Create a Color
// @Description Create a Color
// @Tags Material
// @Accept json
// @produces json
// @Param Request body dto.CreateColorRequest true "Create a Color"
// @Success 201 {object} helper.BaseHttpResponse
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/Material/ [post]
// @Security AuthBearer
func (h *MaterialHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateMaterial, dto.ToMaterialResponse, h.usecase.Create)
}

// UpdateColor godoc
// @Summary Update a Color
// @Description Update a Color
// @Tags Material
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateColorRequest true "Update a Color"
// @Success 200 {object} helper.BaseHttpResponse
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Material/{id} [put]
// @Security AuthBearer
func (h *MaterialHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateMaterial, dto.ToMaterialResponse, h.usecase.Update)
}

// DeleteColor godoc
// @Summary Delete a Color
// @Description Delete a Color
// @Tags Material
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Material/{id} [delete]
// @Security AuthBearer
func (h *MaterialHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetColor godoc
// @Summary Get a Color
// @Description Get a Color
// @Tags Material
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Material/{id} [get]
// @Security AuthBearer
func (h *MaterialHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToMaterialResponse, h.usecase.GetById)
}

// GetMaterial godoc
// @Summary Get Material
// @Description Get Material
// @Tags Material
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/Material/get-by-filter [post]
// @Security AuthBearer
func (h *MaterialHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToMaterialResponse, h.usecase.GetByFilter)
}
