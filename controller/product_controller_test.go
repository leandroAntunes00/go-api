package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-api/dto"
	"go-api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		// Mock Usecase
		mockUsecase := &MockProductUsecase{
			GetProductsFunc: func() ([]model.Product, error) {
				return []model.Product{
					{ID: 1, Name: "Product 1", Price: 10.0},
					{ID: 2, Name: "Product 2", Price: 20.0},
				}, nil
			},
		}

		// Setup router
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		c.Request = req

		// Controller
		productController := NewProductController(mockUsecase)
		productController.GetProducts(c)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		var response []dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, "Product 1", response[0].Name)
		assert.Equal(t, 10.0, response[0].Price)
	})

	t.Run("Error", func(t *testing.T) {
		// Mock Usecase
		mockUsecase := &MockProductUsecase{
			GetProductsFunc: func() ([]model.Product, error) {
				return nil, errors.New("error getting products")
			},
		}

		// Setup router
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		c.Request = req

		// Controller
		productController := NewProductController(mockUsecase)
		productController.GetProducts(c)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		// Mock Usecase
		mockUsecase := &MockProductUsecase{
			CreateProductFunc: func(product model.Product) (model.Product, error) {
				return model.Product{ID: 1, Name: product.Name, Price: product.Price}, nil
			},
		}

		// Request body
		reqBody := dto.CreateProductRequest{
			Name:  "Test Product",
			Price: 25.50,
		}
		jsonBody, _ := json.Marshal(reqBody)

		// Setup router
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		// Controller
		productController := NewProductController(mockUsecase)
		productController.CreateProduct(c)

		// Assertions
		assert.Equal(t, http.StatusCreated, w.Code)
		var response dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Test Product", response.Name)
		assert.Equal(t, 25.50, response.Price)
		assert.Equal(t, 1, response.ID)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		mockUsecase := &MockProductUsecase{}

		// Invalid JSON
		jsonBody := []byte(`{"name": "Test", "price": "invalid"}`)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		productController := NewProductController(mockUsecase)
		productController.CreateProduct(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Usecase Error", func(t *testing.T) {
		mockUsecase := &MockProductUsecase{
			CreateProductFunc: func(product model.Product) (model.Product, error) {
				return model.Product{}, errors.New("database error")
			},
		}

		reqBody := dto.CreateProductRequest{
			Name:  "Test Product",
			Price: 25.50,
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		productController := NewProductController(mockUsecase)
		productController.CreateProduct(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetProductById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockUsecase := &MockProductUsecase{
			GetProductByIdFunc: func(id_product int) (*model.Product, error) {
				return &model.Product{ID: 1, Name: "Test Product", Price: 25.50}, nil
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
		c.Params = gin.Params{{Key: "productId", Value: "1"}}
		c.Request = req

		productController := NewProductController(mockUsecase)
		productController.GetProductById(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var response dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Test Product", response.Name)
		assert.Equal(t, 25.50, response.Price)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		mockUsecase := &MockProductUsecase{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/products/invalid", nil)
		c.Params = gin.Params{{Key: "productId", Value: "invalid"}}
		c.Request = req

		productController := NewProductController(mockUsecase)
		productController.GetProductById(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Product Not Found", func(t *testing.T) {
		mockUsecase := &MockProductUsecase{
			GetProductByIdFunc: func(id_product int) (*model.Product, error) {
				return nil, nil // Product not found
			},
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/products/999", nil)
		c.Params = gin.Params{{Key: "productId", Value: "999"}}
		c.Request = req

		productController := NewProductController(mockUsecase)
		productController.GetProductById(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Empty ID", func(t *testing.T) {
		mockUsecase := &MockProductUsecase{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/products/", nil)
		c.Params = gin.Params{{Key: "productId", Value: ""}}
		c.Request = req

		productController := NewProductController(mockUsecase)
		productController.GetProductById(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
