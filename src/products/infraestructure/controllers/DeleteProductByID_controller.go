package infraestructure

import (
	"demo/src/products/application"
	_"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductByIDController struct {
	dp application.DeleteProductByID
}

func NewDeleteProductByIDController(dp application.DeleteProductByID) *DeleteProductByIDController {
	return &DeleteProductByIDController{dp: dp}
}

func (dp_c *DeleteProductByIDController) Execute(c *gin.Context) {
	// Obtener el parámetro 'id' desde la URL
	idParam := c.Param("id")

	// Si no se pasa un id válido
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	// Convertir el ID de string a int
	idToDelete, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	// Llamar al caso de uso para eliminar el producto
	err = dp_c.dp.Execute(idToDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product", "id": idToDelete})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted product successfully", "id": idToDelete})
}
