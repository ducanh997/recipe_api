package dto

import "recipe_api/model"

type UserDTO struct {
	Username  *string
	Email     *string
	AvatarURL *string
	Age       *int
	ID        uint
	RoleDTOs  []*RoleDTO
}

func NewUserDTO(user *model.User) *UserDTO {
	if user == nil {
		return nil
	}
	roleDTOs := make([]*RoleDTO, 0)
	for _, role := range user.Roles {
		roleDTOs = append(roleDTOs, NewRoleDTO(role))
	}
	return &UserDTO{
		Username: user.Username,
		Email:    user.Email,
		ID:       user.ID,
		RoleDTOs: roleDTOs,
		Age:      user.Age,
	}
}

type UserSearchDTO struct {
	Username *string
	Email    *string
	AgeFrom  *int
	AgeTo    *int
	PageNum  int
	PageSize int
	RoleName *string
}
