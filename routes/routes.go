package routes

import (
	"AgringBackend/controllers"
	"AgringBackend/handlers"
	"AgringBackend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("", controllers.GetOK)
	protected := r.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/vehicles", controllers.GetVehicles)
		protected.POST("/vehicles", controllers.CreateVehicle)
		protected.PUT("/vehicles/:id", controllers.UpdateVehicle)
		protected.DELETE("/vehicles/:id", controllers.DeleteVehicle)

		protected.GET("/vehicle-categories", controllers.GetVehicleCategories)
		protected.POST("/vehicle-categories", controllers.CreateVehicleCategory)
		protected.GET("/vehicle-categories/:id", controllers.GetVehicleCategory)
		protected.PUT("/vehicle-categories/:id", controllers.UpdateVehicleCategory)
		protected.DELETE("/vehicle-categories/:id", controllers.DeleteVehicleCategory)

		protected.GET("/vehicle-statuses", controllers.GetVehicleStatuses)
		protected.POST("/vehicle-statuses", controllers.CreateVehicleStatus)
		protected.GET("/vehicle-statuses/:id", controllers.GetVehicleStatus)
		protected.PUT("/vehicle-statuses/:id", controllers.UpdateVehicleStatus)
		protected.DELETE("/vehicle-statuses/:id", controllers.DeleteVehicleStatus)

		protected.GET("/users", controllers.GetUsers)
		protected.POST("/users", controllers.CreateUser)
		protected.GET("/users/:id", controllers.GetUser)
		protected.PUT("users/:id", controllers.UpdateUser)
		protected.DELETE("users/:id", controllers.DeleteUser)

		protected.GET("/users-roles", controllers.GetUsersRoless)
		protected.POST("/users-roles", controllers.CreateUsersRoles)
		protected.GET("/users-roles/:id", controllers.GetUsersRoles)
		protected.PUT("/users-roles/:id", controllers.UpdateUsersRoles)
		protected.DELETE("/users-roles/:id", controllers.DeleteUsersRoles)

		protected.GET("/roles", controllers.GetRoles)
		protected.POST("/roles", controllers.CreateRole)
		protected.GET("/roles/:id", controllers.GetRole)
		protected.PUT("/roles/:id", controllers.UpdateRole)
		protected.DELETE("/roles/:id", controllers.DeleteRole)

		protected.GET("/fields", controllers.GetFields)
		protected.POST("/fields", controllers.CreateField)
		protected.GET("/fields/:id", controllers.GetField)
		protected.PUT("/fields/:id", controllers.UpdateField)
		protected.DELETE("/fields/:id", controllers.DeleteField)

		protected.GET("/field-workers-model", controllers.GetFieldWorkers)
		protected.POST("/field-workers", controllers.CreateFieldWorkers)
		protected.GET("/field-workers/:id", controllers.GetFieldsWorker)
		protected.PUT("/field-workers/:id", controllers.UpdateFieldWorker)
		protected.DELETE("/field-workers/:id", controllers.DeleteFieldWorker)

		protected.GET("/field-status", controllers.GetFieldStatuses)
		protected.POST("/field-status", controllers.CreateFieldStatus)
		protected.GET("/field-status/:id", controllers.GetFieldStatus)
		protected.PUT("/field-status/:id", controllers.UpdateFieldStatus)
		protected.DELETE("/field-status/:id", controllers.DeleteFieldStatus)

		protected.GET("/crops", controllers.GetCrops)
		protected.POST("/crops", controllers.CreateCrop)
		protected.GET("/crops/:id", controllers.GetCrop)
		protected.PUT("/crops/:id", controllers.UpdateCrop)
		protected.DELETE("/crops/:id", controllers.DeleteCrop)

		protected.GET("/work_types", controllers.GetWorkTypes)
		protected.GET("/work_types/:id", controllers.GetWorkTypeByID)
		protected.POST("/work_types", controllers.CreateWorkType)
		protected.PUT("/work_types/:id", controllers.UpdateWorkType)
		protected.DELETE("/work_types/:id", controllers.DeleteWorkType)
		protected.GET("/field-workers", handlers.GetFieldWorkersHandler)

		protected.Use(gin.Logger())

		protected.GET("/profile", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "username": username})
		})
	}

}
