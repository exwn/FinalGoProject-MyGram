package routers

import (
	"MyGram/config"
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	db := config.StartDB()

	controller := controllers.New(db)

	router := gin.Default()

	router.GET("/users", controller.GetUsers)
	router.POST("/users/login", controller.LoginUser)
	router.POST("/users/register", controller.CreateUsers)
	router.PUT("/users/:userId", controller.UpdateUser)
	// router.DELETE("/users/:userId", controller.DeleteUser)

	UsersGroup := router.Group("/users")
	{
		UsersGroup.Use(middlewares.Authentication())
		UsersGroup.DELETE("/:userId", controller.DeleteUser)
	}

	router.Run()
}
