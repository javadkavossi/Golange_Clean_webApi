package usecase

import (
	"context"

	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/filter"
	model "github.com/javadkavossi/Golange_Clean_webApi/src/domain/model"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/repository"
	"github.com/javadkavossi/Golange_Clean_webApi/src/usecase/dto"
)

type CarModelPropertyUsecase struct {
	base *BaseUsecase[model.CarModelProperty, dto.CreateCarModelProperty, dto.UpdateCarModelProperty, dto.CarModelProperty]
}

func NewCarModelPropertyUsecase(cfg *config.Config, repository repository.CarModelPropertyRepository) *CarModelPropertyUsecase {
	return &CarModelPropertyUsecase{
		base: NewBaseUsecase[model.CarModelProperty, dto.CreateCarModelProperty, dto.UpdateCarModelProperty, dto.CarModelProperty](cfg, repository),
	}
}

// Create
func (u *CarModelPropertyUsecase) Create(ctx context.Context, req dto.CreateCarModelProperty) (dto.CarModelProperty, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelPropertyUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelProperty) (dto.CarModelProperty, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPropertyUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPropertyUsecase) GetById(ctx context.Context, id int) (dto.CarModelProperty, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelPropertyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelProperty], error) {
	return s.base.GetByFilter(ctx, req)
}
