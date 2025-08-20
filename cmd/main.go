package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	_ "go-api/docs" // Importar a documentação Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CRUD GoLang API
// @version 1.0
// @description API REST em Go seguindo Clean Architecture com CRUD de produtos
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /

// @tag.name products
// @tag.description Operações relacionadas a produtos

// @tag.name health
// @tag.description Endpoints de verificação de saúde da API

func main() {
	server := gin.Default()

	// Usar a nova configuração
	dbConfig := db.NewConfig()
	dbConnection, err := db.ConnectDB(dbConfig)
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	//camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	// Swagger documentation endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Ping godoc
	// @Summary Health check
	// @Description Check if the API is running
	// @Tags health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string "API is running"
	// @Router /ping [get]
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Pong"})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/products/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
