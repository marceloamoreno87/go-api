package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)

	user := &entity.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: CreateUser :one INSERT INTO users \( name, email, password \) VALUES \( \$1, \$2, \$3 \) RETURNING id, name, email, password, created_at, updated_at`).
		WithArgs(user.Name, user.Email, user.Password).
		WillReturnRows(rows)

	createdUser, err := ur.CreateUser(user)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)

	user := &entity.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUser :one SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	createdUser, err := ur.GetUser(1)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)

	user := &entity.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUserByEmail :one SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(user.Email).
		WillReturnRows(rows)

	createdUser, err := ur.GetUserByEmail(user.Email)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)

	mock.ExpectExec(`-- name: DeleteUser :exec DELETE FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ur.DeleteUser(1)

	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: UpdateUser :one UPDATE users SET name = \$1, email = \$2, password = \$3 WHERE id = \$4 RETURNING id, name, email, password, created_at, updated_at`).
		WithArgs(user.Name, user.Email, user.Password, user.ID).
		WillReturnRows(rows)

	updatedUser, err := ur.UpdateUser(user, user.ID)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, updatedUser.Name)
	assert.Equal(t, user.Email, updatedUser.Email)
	assert.Equal(t, user.Password, updatedUser.Password)
}
