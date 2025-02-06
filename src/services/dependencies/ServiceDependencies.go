package dependencies

import (
	"database/sql"

	"demo/src/services/application"
	serviceInfra "demo/src/services/infraestructure/controllers"
	serviceRouters "demo/src/services/infraestructure/interfaces/http/routes"
	serviceRepositories "demo/src/services/infraestructure/persistence"

	"github.com/gin-gonic/gin"
)

type ServicesDependencies struct {
	DB *sql.DB
}

func NewServicesDependencies(db *sql.DB) *ServicesDependencies {
	return &ServicesDependencies{DB: db}
}

func (d *ServicesDependencies) Execute(r *gin.Engine) {
	serviceRepository := serviceRepositories.NewServiceRepository(d.DB)
	
	createServiceUseCase := application.NewCreateServiceUseCase(serviceRepository)
	createServiceController := serviceInfra.NewCreateServiceController(createServiceUseCase)

	deleteServiceUseCase := application.NewDeleteServiceUseCase(serviceRepository)
	deleteServiceController := serviceInfra.NewDeleteServiceController(deleteServiceUseCase)

	viewAllServicesUseCase := application.NewViewAllServicesUseCase(serviceRepository)
	getAllServicesController := serviceInfra.NewGetAllServicesController(viewAllServicesUseCase)

	updateServiceUseCase := application.NewEditServiceUseCase(serviceRepository)
	updateServiceController := serviceInfra.NewEditServiceController(updateServiceUseCase)

	
	serviceRouters.AttachServiceRoutes(
		r,
		createServiceController,
		deleteServiceController,
		updateServiceController,
		getAllServicesController,
	)
}
