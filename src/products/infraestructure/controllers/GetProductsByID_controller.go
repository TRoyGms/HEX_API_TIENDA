package infraestructure

import (
	"demo/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIDController struct {
	viewProductByIDUseCase *application.ViewProductByIDUseCase
}

func NewGetProductByIDController(viewProductByIDUseCase *application.ViewProductByIDUseCase) *GetProductByIDController {
	return &GetProductByIDController{viewProductByIDUseCase: viewProductByIDUseCase}
}

func (c *GetProductByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := c.viewProductByIDUseCase.Execute(productID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
