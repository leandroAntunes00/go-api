package repository

import (
	"database/sql"
	"go-api/model"
)


type ProductRepository struct{
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error){
	query := "SELECT id, product_name, price FROM products"
	rows,err := pr.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productsList []model.Product
	var productobj model.Product

	for rows.Next() {
		err = rows.Scan(&productobj.ID, &productobj.Name, &productobj.Price)
		if err != nil {
			return nil, err
		}
		productsList = append(productsList, productobj)
	}
	return productsList, nil
}