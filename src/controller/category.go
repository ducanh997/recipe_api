package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipe_api/src/dto"
)

type CategoryService interface {
	GetCategories(c *gin.Context) ([]*dto.CategoryDTO, error)
}

type CategoryController struct {
	categoryService CategoryService
}

func (t *CategoryController) GetCategories(c *gin.Context) {
	categoryDTOs, err := t.categoryService.GetCategories(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoryDTOs)
}

func NewCategoryController(categoryService CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}
