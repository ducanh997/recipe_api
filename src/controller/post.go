package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipe_api/src/dto"
)

type PostService interface {
	GetPosts(c *gin.Context, searchDTO *dto.PostSearchDTO) ([]*dto.PostDTO, error)
	GetPostByID(c *gin.Context, ID string) (*dto.PostDTO, error)
}

type PostController struct {
	postService PostService
}

func (t *PostController) GetPosts(c *gin.Context) {
	var searchDTO *dto.PostSearchDTO
	err := c.ShouldBind(&searchDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	postDTOs, err := t.postService.GetPosts(c, searchDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, postDTOs)
}

func (t *PostController) GetPostByID(c *gin.Context) {
	postID := c.Param("postID")
	postDTO, err := t.postService.GetPostByID(c, postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, postDTO)
}

func NewPostController(postService PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}
