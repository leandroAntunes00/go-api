package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase interface {
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	GetProductById(id_product int) (*model.Product, error)
}

type productUsecaseImpl struct {
	//repository
	repository repository.ProductRepositoryInterface
}

func NewProductUsecase(repo repository.ProductRepositoryInterface) ProductUsecase {
	return &productUsecaseImpl{
		repository: repo,
	}
}

func (pu *productUsecaseImpl) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *productUsecaseImpl) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *productUsecaseImpl) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
