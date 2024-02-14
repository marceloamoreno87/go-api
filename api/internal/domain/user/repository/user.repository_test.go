package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
	}

	createUserSQL := `-- name: CreateUser :exec INSERT INTO users \( name, email, password, role_id \) VALUES \( \$1, \$2, \$3, \$4 \)`
	mock.ExpectExec(createUserSQL).
		WithArgs(user.Name, user.Email, user.Password, user.RoleID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ur.CreateUser(user)

	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, user.RoleID, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUser :one SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE id = \$1`).
		WithArgs(user.ID).
		WillReturnRows(rows)

	u, err := ur.GetUser(user.ID)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.Password, u.Password)
	assert.Equal(t, user.RoleID, u.RoleID)
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, user.RoleID, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUserByEmail :one SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(user.Email).
		WillReturnRows(rows)

	u, err := ur.GetUserByEmail(user.Email)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.Password, u.Password)
	assert.Equal(t, user.RoleID, u.RoleID)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
	}
	deleteUserSQL := `-- name: DeleteUser :exec DELETE FROM users WHERE id = \$1`

	mock.ExpectExec(deleteUserSQL).
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ur.DeleteUser(user.ID)

	assert.NoError(t, err)

}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)

	user := &entity.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
	}

	updateUserSQL := `-- name: UpdateUser :exec UPDATE users SET name = \$1, email = \$2, password = \$3, role_id = \$4 WHERE id = \$5`
	mock.ExpectExec(updateUserSQL).
		WithArgs(user.Name, user.Email, user.Password, user.RoleID, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ur.UpdateUser(user, 1)

	assert.NoError(t, err)

}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@teste.com",
		Password: "123456",
		RoleID:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, user.RoleID, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUsers :many SELECT id, name, email, password, role_id, created_at, updated_at FROM users ORDER BY id ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(int32(10), int32(0)).
		WillReturnRows(rows)

	users, err := ur.GetUsers(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(users))
	assert.Equal(t, user.Name, users[0].Name)
	assert.Equal(t, user.Email, users[0].Email)
	assert.Equal(t, user.Password, users[0].Password)
	assert.Equal(t, user.RoleID, users[0].RoleID)
}
