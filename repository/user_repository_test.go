package repository

import (
	"database/sql"
	"errors"
	"go-api/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		user := model.User{
			Name:     "Leandro",
			Email:    "leandro@example.com",
			Password: "password123",
		}

		expectedID := 1

		mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id")).
			WithArgs(user.Name, user.Email, user.Password).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

		repo := NewUserRepository(db)
		id, err := repo.CreateUser(user)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_GetUserByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedUser := model.User{
			ID:    1,
			Name:  "Leandro",
			Email: "leandro@example.com",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email FROM users WHERE id = $1")).
			WithArgs(1).
			WillReturnRows(rows)

		repo := NewUserRepository(db)
		user, err := repo.GetUserByID(1)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("User Not Found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email FROM users WHERE id = $1")).
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		repo := NewUserRepository(db)
		user, err := repo.GetUserByID(999)

		assert.NoError(t, err)
		assert.Nil(t, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedUser := model.User{
			ID:       1,
			Name:     "Leandro",
			Email:    "leandro@example.com",
			Password: "password123",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.Password)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password FROM users WHERE email = $1")).
			WithArgs("leandro@example.com").
			WillReturnRows(rows)

		repo := NewUserRepository(db)
		user, err := repo.GetUserByEmail("leandro@example.com")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		user := model.User{
			ID:       1,
			Name:     "Leandro",
			Email:    "leandro@example.com",
			Password: "newpassword",
		}

		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4")).
			WithArgs(user.Name, user.Email, user.Password, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewUserRepository(db)
		err = repo.UpdateUser(user)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = $1")).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewUserRepository(db)
		err = repo.DeleteUser(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_GetUsers(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedUsers := []model.User{
			{ID: 1, Name: "User 1", Email: "user1@example.com"},
			{ID: 2, Name: "User 2", Email: "user2@example.com"},
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(expectedUsers[0].ID, expectedUsers[0].Name, expectedUsers[0].Email).
			AddRow(expectedUsers[1].ID, expectedUsers[1].Name, expectedUsers[1].Email)

		mock.ExpectQuery("SELECT id, name, email FROM users ORDER BY id").
			WillReturnRows(rows)

		repo := NewUserRepository(db)
		users, err := repo.GetUsers()

		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, expectedUsers[0].Name, users[0].Name)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, name, email FROM users ORDER BY id").
			WillReturnError(errors.New("connection failed"))

		repo := NewUserRepository(db)
		users, err := repo.GetUsers()

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Contains(t, err.Error(), "connection failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
