package infraestructure

import (
	"demo/src/services/application"
	"demo/src/services/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateServiceController struct {
	updateServiceUseCase *application.EditServiceUseCase
}

func NewUpdateServiceController(updateServiceUseCase *application.EditServiceUseCase) *UpdateServiceController {
	return &UpdateServiceController{updateServiceUseCase: updateServiceUseCase}
}

func (c *UpdateServiceController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var service entities.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.ID = id

	if err := c.updateServiceUseCase.Execute(&service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, service)
}
