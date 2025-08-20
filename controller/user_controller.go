package controller

import (
	"go-api/dto"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController handles HTTP requests for users
type UserController struct {
	userUsecase usecase.UserUsecase
}

// NewUserController creates a new UserController
func NewUserController(usecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: usecase,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User information"
// @Success 201 {object} dto.UserResponse "User created successfully"
// @Failure 400 {object} model.Response "Bad request - Invalid input data"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /user [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := uc.userUsecase.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, userResponse)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId path int true "User ID" minimum(1)
// @Success 200 {object} dto.UserResponse "User found"
// @Failure 400 {object} model.Response "Bad request - Invalid ID format"
// @Failure 404 {object} model.Response "User not found"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /users/{userId} [get]
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userResponse, err := uc.userUsecase.GetUserByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userResponse == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, userResponse)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param userId path int true "User ID" minimum(1)
// @Param user body dto.UpdateUserRequest true "User information"
// @Success 204 "User updated successfully"
// @Failure 400 {object} model.Response "Bad request - Invalid input data"
// @Failure 404 {object} model.Response "User not found"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /users/{userId} [put]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.userUsecase.UpdateUser(userId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId path int true "User ID" minimum(1)
// @Success 204 "User deleted successfully"
// @Failure 400 {object} model.Response "Bad request - Invalid ID format"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /users/{userId} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = uc.userUsecase.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetUsers godoc
// @Summary List all users
// @Description Get a list of all users in the system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserResponse "List of users"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /users [get]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.userUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
