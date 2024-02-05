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
		RoleId:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, user.RoleId, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: CreateUser :one INSERT INTO users \( name, email, password, role_id \) VALUES \( \$1, \$2, \$3, \$4 \) RETURNING id, name, email, password, role_id, created_at, updated_at`).
		WithArgs(user.Name, user.Email, user.Password, user.RoleId).
		WillReturnRows(rows)

	createdUser, err := ur.CreateUser(user)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.RoleId, createdUser.RoleId)
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
		RoleId:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, user.RoleId, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUser :one SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	createdUser, err := ur.GetUser(1)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.RoleId, createdUser.RoleId)
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
		RoleId:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, user.RoleId, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUserByEmail :one SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(user.Email).
		WillReturnRows(rows)

	createdUser, err := ur.GetUserByEmail(user.Email)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.RoleId, createdUser.RoleId)
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
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleId:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, user.RoleId, time.Now(), time.Now())

	updateUserSQL := `-- name: UpdateUser :one UPDATE users SET name = \$1, email = \$2, password = \$3, role_id = \$4 WHERE id = \$5 RETURNING id, name, email, password, role_id, created_at, updated_at`
	mock.ExpectQuery(updateUserSQL).
		WithArgs(user.Name, user.Email, user.Password, user.RoleId, 1).
		WillReturnRows(rows)

	createdUser, err := ur.UpdateUser(user, 1)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.RoleId, createdUser.RoleId)

}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)

	user := &entity.User{
		Name:     "Test",
		Email:    "test@teste.com",
		Password: "123456",
		RoleId:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(1, user.Name, user.Email, user.Password, user.RoleId, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUsers :many SELECT id, name, email, password, role_id, created_at, updated_at FROM users ORDER BY id ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(int32(10), int32(0)).
		WillReturnRows(rows)

	users, err := ur.GetUsers(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(users))
	assert.Equal(t, user.Name, users[0].Name)
	assert.Equal(t, user.Email, users[0].Email)
	assert.Equal(t, user.Password, users[0].Password)
	assert.Equal(t, user.RoleId, users[0].RoleId)
}
