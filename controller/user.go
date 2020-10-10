package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipe_api/dto"
)

type UserService interface {
	GetUsers(c *gin.Context, searchDTO *dto.UserSearchDTO) ([]*dto.UserDTO, error)
	GetUserByID(c *gin.Context, ID string) (*dto.UserDTO, error)
	CreateUser(c *gin.Context, userDTO *dto.UserDTO) (*dto.UserDTO, error)
	UpdateUser(c *gin.Context, userID string, userDTO *dto.UserDTO) (*dto.UserDTO, error)
}

type UserController struct {
	userService UserService
}

func (t *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	userDTO, err := t.userService.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userDTO)
}

func (t *UserController) GetUsers(c *gin.Context) {
	var searchDTO *dto.UserSearchDTO
	err := c.ShouldBind(&searchDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	userDTOs, err := t.userService.GetUsers(c, searchDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userDTOs)
}

func (t *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("userID")

	var userDTO *dto.UserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user, err := t.userService.UpdateUser(c, userID, userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (t *UserController) CreateUser(c *gin.Context) {
	var userDTO *dto.UserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	userDTO, err = t.userService.CreateUser(c, userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userDTO)
}

func NewUserController(userService UserService) *UserController {
	return &UserController{userService: userService}
}
