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

	api := router.Group("/api")
	{
		api.GET("/users", userController.GetUsers).
			POST("/users", userController.CreateUser).
			GET("/roles", roleController.GetRoles).
			GET("/posts", postController.GetPosts).
			PUT("/users/:userID", userController.UpdateUser)
	}

	router.Run()
}
