package infraestructure

import (
	"demo/src/services/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllServicesController struct {
	getAllServicesUseCase *application.ViewAllServicesUseCase
}

func NewGetAllServicesController(getAllServicesUseCase *application.ViewAllServicesUseCase) *GetAllServicesController {
	return &GetAllServicesController{getAllServicesUseCase: getAllServicesUseCase}
}

func (c *GetAllServicesController) Execute(ctx *gin.Context) {
	services, err := c.getAllServicesUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}
