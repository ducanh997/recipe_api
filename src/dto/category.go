package dto

import "recipe_api/src/model"

type CategoryDTO struct {
	ID   uint
	Name *string
}

func NewCategoryDTO(model *model.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:   model.ID,
		Name: model.Name,
	}
}
