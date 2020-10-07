package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipe_api/dto"
)

type PostService interface {
	GetPosts(c *gin.Context) ([]*dto.PostDTO, error)
}

type PostController struct {
	postService PostService
}

func (t *PostController) GetPosts(c *gin.Context) {
	postDTOs, err := t.postService.GetPosts(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, postDTOs)
}

func NewPostController(postService PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}
