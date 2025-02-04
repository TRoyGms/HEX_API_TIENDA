package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductByIDController struct {
	editProductUseCase *application.EditProduct
}

func NewUpdateProductByIDController(editProductUseCase *application.EditProduct) *UpdateProductByIDController {
	return &UpdateProductByIDController{editProductUseCase: editProductUseCase}
}

func (up_c *UpdateProductByIDController) Execute(c *gin.Context) {
	// Obtener el id del parámetro de la URL
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	// Deserializar el cuerpo de la solicitud
	var updateProduct entities.Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convertir el id de string a int
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	// Usar el ID del producto para actualizar
	updateProduct.Id = id

	// Llamar al caso de uso para editar el producto
	err = up_c.editProductUseCase.Execute(&updateProduct) // Usamos 'editProductUseCase' aquí
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
