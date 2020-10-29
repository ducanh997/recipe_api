package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"recipe_api/src/common/db"
	"recipe_api/src/controller"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Error loading .env file")
	}

	if err := db.Open(); err != nil {
		log.Panic("Error cannot connect to database")
	}

	if err := db.Migrate(); err != nil {
		log.Panic(err.Error())
	}

	router := gin.Default()

	userController := controller.InitUserController()
	roleController := controller.InitRoleController()
	postController := controller.InitPostController()
	categoryController := controller.InitCategoryController()

	api := router.Group("/api/v1")
	{
		user := api.Group("/users")
		{
			user.GET("/", userController.GetUsers)
			user.POST("/", userController.CreateUser)
			user.GET("/:userID", userController.GetUserByID)
			user.PUT("/:userID", userController.UpdateUser)
		}
		role := api.Group("/roles")
		{
			role.GET("/", roleController.GetRoles)
		}
		post := api.Group("/posts")
		{
			post.GET("/", postController.GetPosts)
			post.GET("/:postID", postController.GetPostByID)
		}
		category := api.Group("/categories")
		{
			category.GET("/", categoryController.GetCategories)
		}
	}

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Panic(err.Error())
	}
}
