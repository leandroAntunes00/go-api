package dto

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	// @Description Name of the user
	// @Example "Leandro"
	Name string `json:"name" binding:"required" example:"Leandro"`

	// @Description Email of the user
	// @Example "user@example.com"
	Email string `json:"email" binding:"required,email" example:"user@example.com"`

	// @Description Password of the user
	// @Example "password123"
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	// @Description Name of the user
	// @Example "Leandro"
	Name string `json:"name,omitempty" example:"Leandro"`

	// @Description Email of the user
	// @Example "user@example.com"
	Email string `json:"email,omitempty" example:"user@example.com"`

	// @Description Password of the user
	// @Example "newpassword123"
	Password string `json:"password,omitempty" example:"newpassword123"`
}

// UserResponse represents the response body for user operations
type UserResponse struct {
	// @Description Unique identifier of the user
	// @Example 1
	ID int `json:"id" example:"1"`

	// @Description Name of the user
	// @Example "Leandro"
	Name string `json:"name" example:"Leandro"`

	// @Description Email of the user
	// @Example "user@example.com"
	Email string `json:"email" example:"user@example.com"`
}
