package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"recipe_api/src/common/db"
	"recipe_api/src/dto"
	"recipe_api/src/model"
)

type RoleService struct {
}

func (t *RoleService) GetRoles(c *gin.Context) ([]*dto.RoleDTO, error) {
	var roles []*model.Role
	if err := db.DB.Preload("Users").Find(&roles).Error; err != nil {
		return nil, errors.New("undefined error")
	}

	roleDTOs := make([]*dto.RoleDTO, 0)
	for _, role := range roles {
		roleDTOs = append(roleDTOs, dto.NewRoleDTO(role))
	}
	return roleDTOs, nil
}

func NewRoleService() *RoleService {
	return &RoleService{}
}
