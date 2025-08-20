package controller

import (
	"go-api/model"
)

// MockProductUsecase é um mock do ProductUsecase para testes do controller
type MockProductUsecase struct {
	GetProductsFunc    func() ([]model.Product, error)
	CreateProductFunc  func(product model.Product) (model.Product, error)
	GetProductByIdFunc func(id_product int) (*model.Product, error)
}

func (m *MockProductUsecase) GetProducts() ([]model.Product, error) {
	if m.GetProductsFunc != nil {
		return m.GetProductsFunc()
	}
	return nil, nil
}

func (m *MockProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	if m.CreateProductFunc != nil {
		return m.CreateProductFunc(product)
	}
	return model.Product{}, nil
}

func (m *MockProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	if m.GetProductByIdFunc != nil {
		return m.GetProductByIdFunc(id_product)
	}
	return nil, nil
}
