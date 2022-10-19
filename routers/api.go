package routers

import (
	"MyGram/config"
	"MyGram/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	db := config.StartDB()

	controller := controllers.New(db)

	router := gin.Default()

	router.GET("/users", controller.GetUsers)
	router.POST("/users/register", controller.CreateUsers)
	router.PUT("/users/:userId", controller.UpdateUser)
	router.DELETE("/users/:userId", controller.DeleteUser)

	router.Run()
}
