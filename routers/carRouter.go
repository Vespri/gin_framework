package routers

import (
	"github.com/gin-gonic/gin"

	"gin/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars_all", controllers.GetCar)

	router.GET("/cars/:carID", controllers.GetCar)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
