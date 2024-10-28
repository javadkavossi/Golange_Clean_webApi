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

type CountryHandler struct {
	usecase *usecase.CountryUsecase
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		usecase: usecase.NewCountryUsecase(cfg, dependency.GetCountryRepository(cfg))}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCountryRequest true "Create a country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateUpdateCountry, dto.ToCountryResponse, h.usecase.Create)
}

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCountryRequest true "Update a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	Update(c, dto.ToCreateUpdateCountry, dto.ToCountryResponse, h.usecase.Update)
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCountryResponse, h.usecase.GetById)
}

// GetCountries godoc
// @Summary Get Countries
// @Description Get Countries
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CountryResponse]} "Country response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCountryResponse, h.usecase.GetByFilter)
}
