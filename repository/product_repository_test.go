package repository

import (
	"database/sql"
	"errors"
	"go-api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository_GetProducts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedProducts := []model.Product{
			{ID: 1, Name: "Product 1", Price: 10.0},
			{ID: 2, Name: "Product 2", Price: 20.0},
		}

		rows := sqlmock.NewRows([]string{"id", "product_name", "price"}).
			AddRow(expectedProducts[0].ID, expectedProducts[0].Name, expectedProducts[0].Price).
			AddRow(expectedProducts[1].ID, expectedProducts[1].Name, expectedProducts[1].Price)

		mock.ExpectQuery("SELECT id, product_name, price FROM products").
			WillReturnRows(rows)

		repo := NewProductRepository(db)
		products, err := repo.GetProducts()

		assert.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, expectedProducts[0].Name, products[0].Name)
		assert.Equal(t, expectedProducts[1].Name, products[1].Name)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, product_name, price FROM products").
			WillReturnError(errors.New("connection failed"))

		repo := NewProductRepository(db)
		products, err := repo.GetProducts()

		assert.Error(t, err)
		assert.Nil(t, products)
		assert.Contains(t, err.Error(), "connection failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestProductRepository_CreateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		product := model.Product{
			Name:  "New Product",
			Price: 15.99,
		}

		expectedID := 1

		mock.ExpectPrepare("INSERT INTO products").
			ExpectQuery().
			WithArgs(product.Name, product.Price).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

		repo := NewProductRepository(db)
		id, err := repo.CreateProduct(product)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Prepare Statement Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		product := model.Product{
			Name:  "New Product",
			Price: 15.99,
		}

		mock.ExpectPrepare("INSERT INTO products").
			WillReturnError(errors.New("prepare failed"))

		repo := NewProductRepository(db)
		id, err := repo.CreateProduct(product)

		assert.Error(t, err)
		assert.Equal(t, 0, id)
		assert.Contains(t, err.Error(), "prepare failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Execute Query Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		product := model.Product{
			Name:  "New Product",
			Price: 15.99,
		}

		mock.ExpectPrepare("INSERT INTO products").
			ExpectQuery().
			WithArgs(product.Name, product.Price).
			WillReturnError(errors.New("insert failed"))

		repo := NewProductRepository(db)
		id, err := repo.CreateProduct(product)

		assert.Error(t, err)
		assert.Equal(t, 0, id)
		assert.Contains(t, err.Error(), "insert failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestProductRepository_GetProductById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedProduct := model.Product{
			ID:    1,
			Name:  "Test Product",
			Price: 25.50,
		}

		rows := sqlmock.NewRows([]string{"id", "product_name", "price"}).
			AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Price)

		mock.ExpectPrepare("SELECT \\* FROM products WHERE id = \\$1").
			ExpectQuery().
			WithArgs(1).
			WillReturnRows(rows)

		repo := NewProductRepository(db)
		product, err := repo.GetProductById(1)

		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, expectedProduct.ID, product.ID)
		assert.Equal(t, expectedProduct.Name, product.Name)
		assert.Equal(t, expectedProduct.Price, product.Price)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Product Not Found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare("SELECT \\* FROM products WHERE id = \\$1").
			ExpectQuery().
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		repo := NewProductRepository(db)
		product, err := repo.GetProductById(999)

		assert.NoError(t, err)
		assert.Nil(t, product)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare("SELECT \\* FROM products WHERE id = \\$1").
			ExpectQuery().
			WithArgs(1).
			WillReturnError(errors.New("query failed"))

		repo := NewProductRepository(db)
		product, err := repo.GetProductById(1)

		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Contains(t, err.Error(), "query failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
