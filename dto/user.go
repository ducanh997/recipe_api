package dto

import "recipe_api/model"

type UserDTO struct {
	Username  *string
	Email     *string
	AvatarURL *string
	ID        uint
	RoleDTOs  []*RoleDTO
}

func NewUserDTO(user *model.User) *UserDTO {
	roleDTOs := make([]*RoleDTO, 0)
	for _, role := range user.Roles {
		roleDTOs = append(roleDTOs, NewRoleDTO(role))
	}
	return &UserDTO{
		Username: user.Username,
		Email:    user.Email,
		ID:       user.ID,
		RoleDTOs: roleDTOs,
	}
}

type UserSearchDTO struct {
	Username *string
	Email    *string
	PageNum  int
	PageSize int
}
