package usecase

import (
	"errors"
	"go-api/dto"
	"go-api/model"
	"go-api/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserUsecase defines the contract for the user usecase
type UserUsecase interface {
	CreateUser(user dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUserByID(id int) (*dto.UserResponse, error)
	UpdateUser(id int, user dto.UpdateUserRequest) error
	DeleteUser(id int) error
	GetUsers() ([]dto.UserResponse, error)
}

type userUsecaseImpl struct {
	repository repository.UserRepositoryInterface
}

// NewUserUsecase creates a new instance of UserUsecase
func NewUserUsecase(repo repository.UserRepositoryInterface) UserUsecase {
	return &userUsecaseImpl{
		repository: repo,
	}
}

func (uu *userUsecaseImpl) CreateUser(user dto.CreateUserRequest) (*dto.UserResponse, error) {
	existingUser, err := uu.repository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	id, err := uu.repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (uu *userUsecaseImpl) GetUserByID(id int) (*dto.UserResponse, error) {
	user, err := uu.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (uu *userUsecaseImpl) UpdateUser(id int, user dto.UpdateUserRequest) error {
	existingUser, err := uu.repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUser.Password = string(hashedPassword)
	}

	return uu.repository.UpdateUser(*existingUser)
}

func (uu *userUsecaseImpl) DeleteUser(id int) error {
	return uu.repository.DeleteUser(id)
}

func (uu *userUsecaseImpl) GetUsers() ([]dto.UserResponse, error) {
	users, err := uu.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userResponses, nil
}
