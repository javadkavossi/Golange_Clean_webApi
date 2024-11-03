package dto

import "github.com/javadkavossi/Golange_Clean_webApi/src/usecase/dto"

type CreateMaterialRequest struct {
	Name         string `json:"name" binding:"alpha,min=3,max=15"`
	HexColorCode string `json:"hexCode" binding:"min=7,max=7"`
}

type UpdateMaterialRequest struct {
	Name    string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode,omitempty" binding:"min=7,max=7"`
}

type MaterialResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}

func ToMaterialResponse(from dto.Material) MaterialResponse {
	return MaterialResponse{
		Id:      from.Id,
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}

func ToCreateMaterial(from CreateMaterialRequest) dto.CreateMaterial {
	return dto.CreateMaterial{
		Name:    from.Name,
		HexCode: from.HexColorCode,
	}
}

func ToUpdateMaterial(from CreateMaterialRequest) dto.UpdateMaterial {
	return dto.UpdateMaterial{
		Name:    from.Name,
		HexCode: from.HexColorCode,
	}
}
