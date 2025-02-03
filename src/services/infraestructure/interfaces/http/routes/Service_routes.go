package routers

import (
	infrastructure "demo/src/services/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachServiceRoutes(
	r *gin.Engine,
	createServiceController *infrastructure.CreateServiceController,
	deleteServiceController *infrastructure.DeleteServiceController,
	updateServiceController *infrastructure.EditServiceController,
	getServiceByIDController *infrastructure.GetServiceByIDController,
	getAllServicesController *infrastructure.GetAllServicesController,
) {
	servicesGroup := r.Group("/services")
	{
		servicesGroup.POST("", createServiceController.Execute)       // Crear un nuevo servicio
		servicesGroup.DELETE("/:id", deleteServiceController.Execute) // Eliminar un servicio por ID
		servicesGroup.PUT("/:id", updateServiceController.Execute)    // Actualizar un servicio por ID
		servicesGroup.GET("/:id", getServiceByIDController.Execute)   // Obtener un servicio por ID
		servicesGroup.GET("", getAllServicesController.Execute)       // Obtener todos los servicios
	}
}
