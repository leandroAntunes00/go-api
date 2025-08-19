package repository

import (
	"database/sql"
	"fmt"
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

func (pr * ProductRepository) CreateProduct(product model.Product) (int,error) {
	var id int
	query, err := pr.connection.Prepare(`INSERT INTO products (
		product_name, price
	) VALUES ($1, $2) RETURNING id`)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	query.Close()
	return id,nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error){
	query, err := pr.connection.Prepare(`SELECT * FROM products WHERE id = $1`)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	var product model.Product
	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if(err == sql.ErrNoRows){
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product,nil
}