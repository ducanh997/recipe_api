package dto

import "recipe_api/src/model"

type PostDTO struct {
	Title        *string
	Content      *string
	UserID       *uint
	UserDTO      *UserDTO
	CategoryDTOs []*CategoryDTO
}

func NewPostDTO(post *model.Post) *PostDTO {
	if post == nil {
		return nil
	}

	categoryDTOs := make([]*CategoryDTO, 0)
	for _, category := range post.Categories {
		categoryDTOs = append(categoryDTOs, NewCategoryDTO(category))
	}

	return &PostDTO{
		Title:        post.Title,
		UserID:       post.UserID,
		Content:      post.Content,
		UserDTO:      NewUserDTO(post.User),
		CategoryDTOs: categoryDTOs,
	}
}

type PostSearchDTO struct {
	UserID *uint
}
