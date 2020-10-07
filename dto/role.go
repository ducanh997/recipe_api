package dto

import (
	"recipe_api/model"
)

type RoleDTO struct {
	ID       uint
	Name     *string
	UserDTOs []*UserDTO
}

func NewRoleDTO(role *model.Role) *RoleDTO {
	userDTOs := make([]*UserDTO, 0)
	for _, user := range role.Users {
		userDTOs = append(userDTOs, NewUserDTO(user))
	}
	return &RoleDTO{
		ID:       role.ID,
		Name:     role.Name,
		UserDTOs: userDTOs,
	}
}

func (t *RoleDTO) ToModel() *model.Role {
	return &model.Role{
		Name: t.Name,
	}
}
