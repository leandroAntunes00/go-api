package controller

import (
	"go-api/dto"
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

// MockUserUsecase é um mock do UserUsecase para testes do controller
type MockUserUsecase struct {
	CreateUserFunc  func(user dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUserByIDFunc func(id int) (*dto.UserResponse, error)
	UpdateUserFunc  func(id int, user dto.UpdateUserRequest) error
	DeleteUserFunc  func(id int) error
	GetUsersFunc    func() ([]dto.UserResponse, error)
}

func (m *MockUserUsecase) CreateUser(user dto.CreateUserRequest) (*dto.UserResponse, error) {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(user)
	}
	return nil, nil
}

func (m *MockUserUsecase) GetUserByID(id int) (*dto.UserResponse, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}
	return nil, nil
}

func (m *MockUserUsecase) UpdateUser(id int, user dto.UpdateUserRequest) error {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(id, user)
	}
	return nil
}

func (m *MockUserUsecase) DeleteUser(id int) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(id)
	}
	return nil
}

func (m *MockUserUsecase) GetUsers() ([]dto.UserResponse, error) {
	if m.GetUsersFunc != nil {
		return m.GetUsersFunc()
	}
	return nil, nil
}
