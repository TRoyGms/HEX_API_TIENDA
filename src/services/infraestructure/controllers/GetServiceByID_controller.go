package infraestructure

import (
	"demo/src/services/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetServiceByIDController struct {
	viewServiceByIDUseCase *application.ViewServiceByIDUseCase
}

func NewGetServiceByIDController(viewServiceByIDUseCase *application.ViewServiceByIDUseCase) *GetServiceByIDController {
	return &GetServiceByIDController{viewServiceByIDUseCase: viewServiceByIDUseCase}
}

func (c *GetServiceByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	serviceID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	service, err := c.viewServiceByIDUseCase.Execute(serviceID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	ctx.JSON(http.StatusOK, service)
}
