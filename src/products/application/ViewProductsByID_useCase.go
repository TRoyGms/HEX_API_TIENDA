package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
	"errors"
)

type ViewProductByIDUseCase struct {
	productRepository domain.IProduct
}

func NewViewProductByIDUseCase(productRepository domain.IProduct) *ViewProductByIDUseCase {
	return &ViewProductByIDUseCase{productRepository: productRepository}
}

func (u *ViewProductByIDUseCase) Execute(productID int) (*entities.Product, error) {
	product, err := u.productRepository.GetByID(productID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
