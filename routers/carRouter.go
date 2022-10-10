package routers

import (
	"github.com/gin-gonic/gin"

	"gin/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	return router
}
