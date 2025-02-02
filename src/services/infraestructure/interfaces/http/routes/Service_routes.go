package routers

import (
    infrastructure "demo/src/services/infraestructure/controllers"

    "github.com/gin-gonic/gin"
)

func AttachServiceRoutes(
	r *gin.Engine,
	createServiceController *infrastructure.CreateServiceController,
	// deleteServiceController *infrastructure.DeleteServiceByIDController,
	// editServiceController *infrastructure.EditServiceByIDController,
	// getServiceController *infrastructure.GetServiceByIDController,
	// getAllServicesController *infrastructure.GetAllServicesController,
) {
	servicesGroup := r.Group("/services")
		servicesGroup.POST("", createServiceController.Execute)
		// servicesGroup.DELETE("/:id", deleteServiceController.Execute)
		// servicesGroup.PUT("/:id", editServiceController.Execute)
		// servicesGroup.GET("/:id", getServiceController.Execute)
		// servicesGroup.GET("", getAllServicesController.Execute)
}

