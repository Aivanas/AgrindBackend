package routes

import (
	"AgringBackend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/vehicles", controllers.GetVehicles)
	r.POST("/vehicles", controllers.CreateVehicle)
	r.PUT("/vehicles/:id", controllers.UpdateVehicle)
	r.DELETE("/vehicles/:id", controllers.DeleteVehicle)
	r.GET("", controllers.GetOK)
}
