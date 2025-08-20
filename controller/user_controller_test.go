package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-api/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			CreateUserFunc: func(user dto.CreateUserRequest) (*dto.UserResponse, error) {
				return &dto.UserResponse{ID: 1, Name: user.Name, Email: user.Email}, nil
			},
		}

		reqBody := dto.CreateUserRequest{
			Name:     "Leandro",
			Email:    "leandro@example.com",
			Password: "password123",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.CreateUser(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response dto.UserResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Leandro", response.Name)
		assert.Equal(t, "leandro@example.com", response.Email)
		assert.Equal(t, 1, response.ID)
	})
}

func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			GetUserByIDFunc: func(id int) (*dto.UserResponse, error) {
				return &dto.UserResponse{ID: 1, Name: "Leandro", Email: "leandro@example.com"}, nil
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		c.Params = gin.Params{{Key: "userId", Value: "1"}}
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.GetUserByID(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var response dto.UserResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Leandro", response.Name)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			GetUserByIDFunc: func(id int) (*dto.UserResponse, error) {
				return nil, nil
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		c.Params = gin.Params{{Key: "userId", Value: "1"}}
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.GetUserByID(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			UpdateUserFunc: func(id int, user dto.UpdateUserRequest) error {
				return nil
			},
		}

		reqBody := dto.UpdateUserRequest{
			Name: "Leandro Updated",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		c.Params = gin.Params{{Key: "userId", Value: "1"}}
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.UpdateUser(c)

		assert.Equal(t, http.StatusNoContent, c.Writer.Status())
	})
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			DeleteUserFunc: func(id int) error {
				return nil
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		c.Params = gin.Params{{Key: "userId", Value: "1"}}
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.DeleteUser(c)

		assert.Equal(t, http.StatusNoContent, c.Writer.Status())
	})

	t.Run("Error", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			DeleteUserFunc: func(id int) error {
				return errors.New("delete failed")
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		c.Params = gin.Params{{Key: "userId", Value: "1"}}
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.DeleteUser(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockUsecase := &MockUserUsecase{
			GetUsersFunc: func() ([]dto.UserResponse, error) {
				return []dto.UserResponse{
					{ID: 1, Name: "User 1", Email: "user1@example.com"},
					{ID: 2, Name: "User 2", Email: "user2@example.com"},
				}, nil
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		c.Request = req

		userController := NewUserController(mockUsecase)
		userController.GetUsers(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var response []dto.UserResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
	})
}
