package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"recipe_api/common"
	"recipe_api/common/db"
	"recipe_api/dto"
	"recipe_api/model"
)

type UserService struct {
}

func (t *UserService) CreateUser(c *gin.Context, userDTO *dto.UserDTO) (*dto.UserDTO, error) {
	role := &model.Role{}
	if err := db.DB.First(&role, model.Role{Name: common.PStr("Member")}).Error; err != nil {
		return nil, errors.New("undefined error")
	}

	user := &model.User{
		Username:  userDTO.Username,
		Email:     userDTO.Email,
		AvatarURL: nil,
		Roles:     []*model.Role{role},
	}

	if err := db.DB.Create(user).Error; err != nil {
		return nil, errors.New("undefined error")
	}
	return dto.NewUserDTO(user), nil
}

func (t *UserService) UpdateUser(c *gin.Context, userID string, userDTO *dto.UserDTO) (*dto.UserDTO, error) {
	user := &model.User{}
	if err := db.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("undefined error")
	}

	user.Username = userDTO.Username
	user.AvatarURL = userDTO.AvatarURL
	user.Email = userDTO.Email

	if err := db.DB.Save(&user).Error; err != nil {
		return nil, errors.New("undefined error")
	}
	return userDTO, nil
}

func (t *UserService) GetUsers(c *gin.Context, searchDTO *dto.UserSearchDTO) ([]*dto.UserDTO, error) {
	if searchDTO == nil {
		searchDTO = &dto.UserSearchDTO{}
	}

	var users []*model.User
	query := db.DB.Debug().Preload("Roles")

	if searchDTO.Email != nil {
		query = query.Where(model.User{Email: searchDTO.Email})
	}
	if searchDTO.Username != nil {
		query = query.Where(model.User{Username: searchDTO.Username})
	}

	query = query.Order("ID desc")
	query.Scopes(db.Paginate(searchDTO.PageNum, searchDTO.PageSize))

	if err := query.Find(&users).Error; err != nil {
		return nil, errors.New("undefined error")
	}

	userDTOs := make([]*dto.UserDTO, 0)
	for _, user := range users {
		userDTOs = append(userDTOs, dto.NewUserDTO(user))
	}

	return userDTOs, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
