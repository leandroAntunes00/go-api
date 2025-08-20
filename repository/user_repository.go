package repository

import (
	"database/sql"
	"go-api/model"
)

// UserRepositoryInterface defines the contract for the user repository
type UserRepositoryInterface interface {
	CreateUser(user model.User) (int, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
	GetUsers() ([]model.User, error)
}

type UserRepository struct {
	connection *sql.DB
}

// Ensure UserRepository implements UserRepositoryInterface
var _ UserRepositoryInterface = (*UserRepository)(nil)

func NewUserRepository(connection *sql.DB) UserRepositoryInterface {
	return &UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	err := ur.connection.QueryRow(`INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := ur.connection.QueryRow(`SELECT id, name, email FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := ur.connection.QueryRow(`SELECT id, name, email, password FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(user model.User) error {
	_, err := ur.connection.Exec(`UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`, user.Name, user.Email, user.Password, user.ID)
	return err
}

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.connection.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	rows, err := ur.connection.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
