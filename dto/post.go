package dto

import "recipe_api/model"

type PostDTO struct {
	Title   *string
	UserID  *uint
	UserDTO *UserDTO
}

func NewPostDTO(post *model.Post) *PostDTO {
	if post == nil {
		return nil
	}
	return &PostDTO{
		Title:   post.Title,
		UserID:  post.UserID,
		UserDTO: NewUserDTO(post.User),
	}
}
