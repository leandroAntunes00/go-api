package controller

import (
	"go-api/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductController struct{
	//Usecase
	productUsecase usecase.Product_usecase
}


func NewProductController(usecase usecase.Product_usecase) *ProductController {
	return &ProductController{
		//Usecase
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}
