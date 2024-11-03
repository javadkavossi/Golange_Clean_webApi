package usecase

import (
	"context"

	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/filter"
	model "github.com/javadkavossi/Golange_Clean_webApi/src/domain/model"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/repository"
	"github.com/javadkavossi/Golange_Clean_webApi/src/usecase/dto"
)

type MaterialUsecase struct {
	base *BaseUsecase[model.Material, dto.CreateMaterial, dto.UpdateMaterial, dto.Material]
}

func NewMaterialUsecase(cfg *config.Config, repository repository.MaterialRepository) *MaterialUsecase {
	return &MaterialUsecase{
		base: NewBaseUsecase[model.Material, dto.CreateMaterial, dto.UpdateMaterial, dto.Material](cfg, repository),
	}
}

// Create
func (u *MaterialUsecase) Create(ctx context.Context, req dto.CreateMaterial) (dto.Material, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *MaterialUsecase) Update(ctx context.Context, id int, req dto.UpdateMaterial) (dto.Material, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *MaterialUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *MaterialUsecase) GetById(ctx context.Context, id int) (dto.Material, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *MaterialUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Material], error) {
	return s.base.GetByFilter(ctx, req)
}
