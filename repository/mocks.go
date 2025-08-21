package repository

import (
	"go-api/model"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for the UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

// CreateUser mocks the CreateUser method
func (m *MockUserRepository) CreateUser(user model.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

// GetUserByID mocks the GetUserByID method
func (m *MockUserRepository) GetUserByID(id int) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

// GetUserByEmail mocks the GetUserByEmail method
func (m *MockUserRepository) GetUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

// UpdateUser mocks the UpdateUser method
func (m *MockUserRepository) UpdateUser(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// DeleteUser mocks the DeleteUser method
func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetUsers mocks the GetUsers method
func (m *MockUserRepository) GetUsers() ([]model.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), args.Error(1)
}
