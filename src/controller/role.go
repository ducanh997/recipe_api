package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipe_api/src/dto"
)

type RoleService interface {
	GetRoles(c *gin.Context) ([]*dto.RoleDTO, error)
}

type RoleController struct {
	roleService RoleService
}

func (t *RoleController) GetRoles(c *gin.Context) {
	roles, err := t.roleService.GetRoles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func NewRoleController(roleService RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}
