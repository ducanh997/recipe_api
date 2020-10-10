//+build wireinject

package controller

import (
	"github.com/google/wire"
	"recipe_api/service"
)

func InitUserController() *UserController {
	wire.Build(
		NewUserController,
		service.NewUserService,
		wire.Bind(new(UserService), new(*service.UserService)),
	)

	return &UserController{}
}

func InitRoleController() *RoleController {
	wire.Build(
		NewRoleController,
		service.NewRoleService,
		wire.Bind(new(RoleService), new(*service.RoleService)),
	)

	return &RoleController{}
}

func InitPostController() *PostController {
	wire.Build(
		NewPostController,
		service.NewPostService,
		wire.Bind(new(PostService), new(*service.PostService)),
	)

	return &PostController{}
}

func InitAuthController() *AuthController {
	wire.Build(
		NewAuthController,
	)
	return &AuthController{}
}
