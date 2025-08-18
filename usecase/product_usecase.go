package usecase

import (
	"go-api/model"
	"go-api/repository"
)


type Product_usecase struct{
	//repository
	repository repository.ProductRepository
}


func NewProductUsecase(repo repository.ProductRepository) Product_usecase {
	return Product_usecase{
		repository: repo,
	}
}

func (pu *Product_usecase) GetProducts() ([]model.Product, error){
	return pu.repository.GetProducts()
}