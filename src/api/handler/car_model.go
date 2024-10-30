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

type CarModelHandler struct {
	usecase *usecase.CarModelUsecase
}

// CarModelPagedList struct for Swagger documentation
type CarModelPagedList struct {
    Items      []dto.CarModelResponse `json:"items"`
    TotalCount int                    `json:"total_count"`
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		usecase: usecase.NewCarModelUsecase(cfg, dependency.GetCarModelRepository(cfg)),
	}
}

// CreateCarModel godoc
// @Summary Create a CarModel
// @Description Create a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelRequest true "Create a CarModel"

// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModel, dto.ToCarModelResponse, h.usecase.Create)
}

// UpdateCarModel godoc
// @Summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelRequest true "Update a CarModel"
// @Success 200 {object} helper.BaseHttpResponse"CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModel, dto.ToCarModelResponse, h.usecase.Update)
}

// DeleteCarModel godoc
// @Summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModel godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelResponse, h.usecase.GetById)
}

// GetCarModels godoc
// @Summary Get CarModels
// @Description Get CarModels
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"

// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelResponse, h.usecase.GetByFilter)
}
