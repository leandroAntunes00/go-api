package usecase

import (
	"go-api/model"
)

// MockProductRepository é um mock do ProductRepository para testes do usecase
type MockProductRepository struct {
	GetProductsFunc    func() ([]model.Product, error)
	CreateProductFunc  func(product model.Product) (int, error)
	GetProductByIdFunc func(id_product int) (*model.Product, error)
}

func (m *MockProductRepository) GetProducts() ([]model.Product, error) {
	if m.GetProductsFunc != nil {
		return m.GetProductsFunc()
	}
	return nil, nil
}

func (m *MockProductRepository) CreateProduct(product model.Product) (int, error) {
	if m.CreateProductFunc != nil {
		return m.CreateProductFunc(product)
	}
	return 0, nil
}

func (m *MockProductRepository) GetProductById(id_product int) (*model.Product, error) {
	if m.GetProductByIdFunc != nil {
		return m.GetProductByIdFunc(id_product)
	}
	return nil, nil
}
