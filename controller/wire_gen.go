// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controller

import (
	"recipe_api/service"
)

// Injectors from wire.go:

func InitUserController() *UserController {
	userService := service.NewUserService()
	userController := NewUserController(userService)
	return userController
}

func InitRoleController() *RoleController {
	roleService := service.NewRoleService()
	roleController := NewRoleController(roleService)
	return roleController
}

func InitPostController() *PostController {
	postService := service.NewPostService()
	postController := NewPostController(postService)
	return postController
}
