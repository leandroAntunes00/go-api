package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type Product_usecase struct {
	//repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) Product_usecase {
	return Product_usecase{
		repository: repo,
	}
}

func (pu *Product_usecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *Product_usecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *Product_usecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
