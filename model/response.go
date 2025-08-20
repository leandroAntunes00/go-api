package model

// Response represents a generic API response
type Response struct {
	// @Description Response message
	// @Example "Operation completed successfully"
	Message string `json:"message" example:"Operation completed successfully"`
}
