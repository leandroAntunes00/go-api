package main

import (
	"go-api/controller"
	"go-api/db"
	_ "go-api/docs" // Importar a documentação Swagger
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CRUD GoLang API
// @version 1.0
// @description API REST em Go seguindo Clean Architecture com CRUD de produtos e usuários
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

// @tag.name users
// @tag.description Operações relacionadas a usuários

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

	// Product
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	// User
	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsecase := usecase.NewUserUsecase(UserRepository)
	UserController := controller.NewUserController(UserUsecase)

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

	// Product routes
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/products/:productId", ProductController.GetProductById)

	// User routes
	server.POST("/user", UserController.CreateUser)
	server.GET("/users/:userId", UserController.GetUserByID)
	server.PUT("/users/:userId", UserController.UpdateUser)
	server.DELETE("/users/:userId", UserController.DeleteUser)
	server.GET("/users", UserController.GetUsers)

	// Login route
	server.POST("/login", UserController.Login)

	server.Run(":8000")
}
