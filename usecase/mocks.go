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

// MockUserRepository é um mock do UserRepository para testes do usecase
type MockUserRepository struct {
	CreateUserFunc     func(user model.User) (int, error)
	GetUserByIDFunc    func(id int) (*model.User, error)
	GetUserByEmailFunc func(email string) (*model.User, error)
	UpdateUserFunc     func(user model.User) error
	DeleteUserFunc     func(id int) error
	GetUsersFunc       func() ([]model.User, error)
}

func (m *MockUserRepository) CreateUser(user model.User) (int, error) {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(user)
	}
	return 0, nil
}

func (m *MockUserRepository) GetUserByID(id int) (*model.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}
	return nil, nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (*model.User, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(email)
	}
	return nil, nil
}

func (m *MockUserRepository) UpdateUser(user model.User) error {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(user)
	}
	return nil
}

func (m *MockUserRepository) DeleteUser(id int) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(id)
	}
	return nil
}

func (m *MockUserRepository) GetUsers() ([]model.User, error) {
	if m.GetUsersFunc != nil {
		return m.GetUsersFunc()
	}
	return nil, nil
}
