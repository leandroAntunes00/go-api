package dto

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	// @Description Name of the product
	// @Example "iPhone 15"
	Name string `json:"name" binding:"required" example:"iPhone 15"`
	
	// @Description Price of the product
	// @Example 999.99
	Price float64 `json:"price" binding:"required,min=0" example:"999.99"`
}

// ProductResponse represents the response body for product operations
type ProductResponse struct {
	// @Description Unique identifier of the product
	// @Example 1
	ID int `json:"id" example:"1"`
	
	// @Description Name of the product
	// @Example "iPhone 15"
	Name string `json:"name" example:"iPhone 15"`
	
	// @Description Price of the product
	// @Example 999.99
	Price float64 `json:"price" example:"999.99"`
}
