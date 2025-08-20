package controller

import (
	"go-api/dto"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController handles HTTP requests for products
type ProductController struct {
	//Usecase
	productUsecase usecase.ProductUsecase
}

// NewProductController creates a new ProductController
func NewProductController(usecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		//Usecase
		productUsecase: usecase,
	}
}

// GetProducts godoc
// @Summary List all products
// @Description Get a list of all products in the system
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} dto.ProductResponse "List of products"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /products [get]
func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, toProductResponse(product))
	}

	ctx.JSON(http.StatusOK, productResponses)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product information"
// @Success 201 {object} dto.ProductResponse "Product created successfully"
// @Failure 400 {object} model.Response "Bad request - Invalid input data"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /product [post]
func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := toProductModel(req)

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, toProductResponse(insertedProduct))
}

// GetProductById godoc
// @Summary Get product by ID
// @Description Get a specific product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param productId path int true "Product ID" minimum(1)
// @Success 200 {object} dto.ProductResponse "Product found"
// @Failure 400 {object} model.Response "Bad request - Invalid ID format"
// @Failure 404 {object} model.Response "Product not found"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /products/{productId} [get]
func (p *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "id do produto precisa ser numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		response := model.Response{
			Message: "produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, toProductResponse(*product))
}

// --- Helper Functions ---

func toProductModel(req dto.CreateProductRequest) model.Product {
	return model.Product{
		Name:  req.Name,
		Price: req.Price,
	}
}

func toProductResponse(product model.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}
