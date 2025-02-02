package infraestructure

import (
	"demo/src/services/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteServiceController struct {
	deleteServiceUseCase *application.DeleteServiceUseCase
}

func NewDeleteServiceController(deleteServiceUseCase *application.DeleteServiceUseCase) *DeleteServiceController {
	return &DeleteServiceController{deleteServiceUseCase: deleteServiceUseCase}
}

func (c *DeleteServiceController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	if err := c.deleteServiceUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}
