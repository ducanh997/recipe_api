package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"recipe_api/common/db"
	"recipe_api/dto"
	"recipe_api/model"
)

type PostService struct {
}

func (t *PostService) GetPosts(c *gin.Context) ([]*dto.PostDTO, error) {
	var posts []*model.Post
	if err := db.DB.Preload("User").Find(&posts).Error; err != nil {
		return nil, errors.New("undefined error")
	}

	postDTOs := make([]*dto.PostDTO, 0)
	for _, post := range posts {
		postDTOs = append(postDTOs, dto.NewPostDTO(post))
	}
	return postDTOs, nil
}

func NewPostService() *PostService {
	return &PostService{}
}
