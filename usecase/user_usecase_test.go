package usecase

import (
	"errors"
	"go-api/dto"
	"go-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		createUserRequest := dto.CreateUserRequest{
			Name:     "Leandro",
			Email:    "leandro@example.com",
			Password: "password123",
		}

		mockRepo := &MockUserRepository{
			GetUserByEmailFunc: func(email string) (*model.User, error) {
				return nil, nil
			},
			CreateUserFunc: func(user model.User) (int, error) {
				return 1, nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		userResponse, err := usecase.CreateUser(createUserRequest)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, 1, userResponse.ID)
		assert.Equal(t, createUserRequest.Name, userResponse.Name)
		assert.Equal(t, createUserRequest.Email, userResponse.Email)
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		createUserRequest := dto.CreateUserRequest{
			Name:     "Leandro",
			Email:    "leandro@example.com",
			Password: "password123",
		}

		mockRepo := &MockUserRepository{
			GetUserByEmailFunc: func(email string) (*model.User, error) {
				return &model.User{}, nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		userResponse, err := usecase.CreateUser(createUserRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Equal(t, "user with this email already exists", err.Error())
	})
}

func TestUserUsecase_GetUserByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedUser := &model.User{
			ID:    1,
			Name:  "Leandro",
			Email: "leandro@example.com",
		}

		mockRepo := &MockUserRepository{
			GetUserByIDFunc: func(id int) (*model.User, error) {
				return expectedUser, nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		userResponse, err := usecase.GetUserByID(1)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, expectedUser.ID, userResponse.ID)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo := &MockUserRepository{
			GetUserByIDFunc: func(id int) (*model.User, error) {
				return nil, nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		userResponse, err := usecase.GetUserByID(1)

		assert.NoError(t, err)
		assert.Nil(t, userResponse)
	})
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	existingUser := &model.User{
		ID:       1,
		Name:     "Leandro",
		Email:    "leandro@example.com",
		Password: string(hashedPassword),
	}

	t.Run("Success", func(t *testing.T) {
		updateUserRequest := dto.UpdateUserRequest{
			Name: "Leandro Updated",
		}

		mockRepo := &MockUserRepository{
			GetUserByIDFunc: func(id int) (*model.User, error) {
				return existingUser, nil
			},
			UpdateUserFunc: func(user model.User) error {
				return nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		err := usecase.UpdateUser(1, updateUserRequest)

		assert.NoError(t, err)
	})

	t.Run("User Not Found", func(t *testing.T) {
		updateUserRequest := dto.UpdateUserRequest{
			Name: "Leandro Updated",
		}

		mockRepo := &MockUserRepository{
			GetUserByIDFunc: func(id int) (*model.User, error) {
				return nil, nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		err := usecase.UpdateUser(1, updateUserRequest)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockUserRepository{
			DeleteUserFunc: func(id int) error {
				return nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		err := usecase.DeleteUser(1)

		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := &MockUserRepository{
			DeleteUserFunc: func(id int) error {
				return errors.New("delete failed")
			},
		}

		usecase := NewUserUsecase(mockRepo)
		err := usecase.DeleteUser(1)

		assert.Error(t, err)
		assert.Equal(t, "delete failed", err.Error())
	})
}

func TestUserUsecase_GetUsers(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedUsers := []model.User{
			{ID: 1, Name: "User 1", Email: "user1@example.com"},
			{ID: 2, Name: "User 2", Email: "user2@example.com"},
		}

		mockRepo := &MockUserRepository{
			GetUsersFunc: func() ([]model.User, error) {
				return expectedUsers, nil
			},
		}

		usecase := NewUserUsecase(mockRepo)
		userResponses, err := usecase.GetUsers()

		assert.NoError(t, err)
		assert.Len(t, userResponses, 2)
		assert.Equal(t, expectedUsers[0].Name, userResponses[0].Name)
	})
}
