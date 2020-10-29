package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"recipe_api/src/common/db"
	"recipe_api/src/dto"
	"recipe_api/src/model"
)

type CategoryService struct {
}

func (t *CategoryService) GetCategories(c *gin.Context) ([]*dto.CategoryDTO, error) {
	var categories []*model.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		return nil, errors.New("undefined error")
	}

	categoryDTOs := make([]*dto.CategoryDTO, 0)
	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, dto.NewCategoryDTO(category))
	}
	fmt.Print("Hello")
	return categoryDTOs, nil
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}
