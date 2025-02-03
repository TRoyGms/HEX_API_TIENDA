package infraestructure

import (
	"demo/src/services/application"
	"demo/src/services/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditServiceController struct {
	editServiceUseCase *application.EditServiceUseCase
}

func NewEditServiceController(editServiceUseCase *application.EditServiceUseCase) *EditServiceController {
	return &EditServiceController{editServiceUseCase: editServiceUseCase}
}

func (c *EditServiceController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	serviceID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var updatedService entities.Service
	if err := ctx.ShouldBindJSON(&updatedService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.editServiceUseCase.Execute(serviceID, &updatedService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Service updated successfully"})
}
