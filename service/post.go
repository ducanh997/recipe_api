package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (t *PostService) GetPostByID(c *gin.Context, ID string) (*dto.PostDTO, error) {
	post := &model.Post{}
	if err := db.DB.First(&post, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, errors.New("undefined error")
	}
	return dto.NewPostDTO(post), nil
}

func NewPostService() *PostService {
	return &PostService{}
}
