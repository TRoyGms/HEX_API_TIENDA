package application

import (
	"demo/src/products/domain/entities"
	repositories "demo/src/products/infraestructure/persistence"
)

type EditProduct struct {
	db repositories.ProductRepository
}

func NewEditProduct(db repositories.ProductRepository) *EditProduct {
	return &EditProduct{db: db}
}

func (cp *EditProduct) Execute(updatedProduct *entities.Product) error {
	// Se utiliza el `updatedProduct` para actualizar el producto sin necesidad de pasar el `productID`
	err := cp.db.EditById(int(updatedProduct.Id), updatedProduct)
	if err != nil {
		return err
	}
	return nil
}