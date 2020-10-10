package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"recipe_api/common/db"
	"recipe_api/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	if err := db.Open(); err != nil {
		log.Panic("Error cannot connect to database")
	}

	if err = db.Migrate(); err != nil {
		log.Panic(err.Error())
	}

	router := gin.Default()

	userController := controller.InitUserController()
	roleController := controller.InitRoleController()
	postController := controller.InitPostController()

	api := router.Group("/api/v1")
	{
		user := api.Group("/users")
		{
			user.GET("/", userController.GetUsers)
			user.POST("/", userController.CreateUser)
			user.PUT("/:userID", userController.UpdateUser)
		}
		role := api.Group("/roles")
		{
			role.GET("/", roleController.GetRoles)
		}
		post := api.Group("/posts")
		{
			post.GET("/", postController.GetPosts)
		}
	}

	router.Run()
}
