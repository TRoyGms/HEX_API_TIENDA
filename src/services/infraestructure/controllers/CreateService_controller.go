package infraestructure

import (
	"demo/src/services/application"
	"demo/src/services/domain/entities"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateServiceController struct {
	createServiceUseCase *application.CreateServiceUseCase
}

func NewCreateServiceController(createServiceUseCase *application.CreateServiceUseCase) *CreateServiceController {
	return &CreateServiceController{createServiceUseCase: createServiceUseCase}
}

func (c *CreateServiceController) Execute(ctx *gin.Context) {
	var service entities.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.createServiceUseCase.Execute(&service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, service)
}
