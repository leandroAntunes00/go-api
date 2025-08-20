package usecase

import (
	"errors"
	"go-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductUsecase_GetProducts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedProducts := []model.Product{
			{ID: 1, Name: "Product 1", Price: 10.0},
			{ID: 2, Name: "Product 2", Price: 20.0},
		}

		mockRepo := &MockProductRepository{
			GetProductsFunc: func() ([]model.Product, error) {
				return expectedProducts, nil
			},
		}

		usecase := NewProductUsecase(mockRepo)
		products, err := usecase.GetProducts()

		assert.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, expectedProducts[0].Name, products[0].Name)
		assert.Equal(t, expectedProducts[1].Name, products[1].Name)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := &MockProductRepository{
			GetProductsFunc: func() ([]model.Product, error) {
				return nil, errors.New("database connection failed")
			},
		}

		usecase := NewProductUsecase(mockRepo)
		products, err := usecase.GetProducts()

		assert.Error(t, err)
		assert.Nil(t, products)
		assert.Contains(t, err.Error(), "database connection failed")
	})
}

func TestProductUsecase_CreateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		productToCreate := model.Product{
			Name:  "New Product",
			Price: 15.99,
		}

		mockRepo := &MockProductRepository{
			CreateProductFunc: func(product model.Product) (int, error) {
				return 1, nil
			},
		}

		usecase := NewProductUsecase(mockRepo)
		createdProduct, err := usecase.CreateProduct(productToCreate)

		assert.NoError(t, err)
		assert.Equal(t, 1, createdProduct.ID)
		assert.Equal(t, productToCreate.Name, createdProduct.Name)
		assert.Equal(t, productToCreate.Price, createdProduct.Price)
	})

	t.Run("Repository Error", func(t *testing.T) {
		productToCreate := model.Product{
			Name:  "New Product",
			Price: 15.99,
		}

		mockRepo := &MockProductRepository{
			CreateProductFunc: func(product model.Product) (int, error) {
				return 0, errors.New("insert failed")
			},
		}

		usecase := NewProductUsecase(mockRepo)
		createdProduct, err := usecase.CreateProduct(productToCreate)

		assert.Error(t, err)
		assert.Equal(t, model.Product{}, createdProduct)
		assert.Contains(t, err.Error(), "insert failed")
	})
}

func TestProductUsecase_GetProductById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedProduct := &model.Product{
			ID:    1,
			Name:  "Test Product",
			Price: 25.50,
		}

		mockRepo := &MockProductRepository{
			GetProductByIdFunc: func(id_product int) (*model.Product, error) {
				return expectedProduct, nil
			},
		}

		usecase := NewProductUsecase(mockRepo)
		product, err := usecase.GetProductById(1)

		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, expectedProduct.ID, product.ID)
		assert.Equal(t, expectedProduct.Name, product.Name)
		assert.Equal(t, expectedProduct.Price, product.Price)
	})

	t.Run("Product Not Found", func(t *testing.T) {
		mockRepo := &MockProductRepository{
			GetProductByIdFunc: func(id_product int) (*model.Product, error) {
				return nil, nil
			},
		}

		usecase := NewProductUsecase(mockRepo)
		product, err := usecase.GetProductById(999)

		assert.NoError(t, err)
		assert.Nil(t, product)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := &MockProductRepository{
			GetProductByIdFunc: func(id_product int) (*model.Product, error) {
				return nil, errors.New("query failed")
			},
		}

		usecase := NewProductUsecase(mockRepo)
		product, err := usecase.GetProductById(1)

		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Contains(t, err.Error(), "query failed")
	})
}
