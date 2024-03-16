package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/config"
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
	config.NewDatabaseMock(db)
	repo := repository.NewUserRepository(config.DbMock)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
		AvatarID: 1,
	}

	createUserSQL := `-- name: CreateUser :exec INSERT INTO users \( name, email, password, role_id, avatar_id \) VALUES \( \$1, \$2, \$3, \$4, \$5 \)`
	mock.ExpectBegin()
	mock.ExpectExec(createUserSQL).
		WithArgs(user.Name, user.Email, user.Password, user.RoleID, user.AvatarID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = repo.Begin()
	assert.NoError(t, err)
	newuser := repo.CreateUser(user)
	assert.NoError(t, newuser)
	err = repo.Commit()
	assert.NoError(t, err)

}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewUserRepository(config.DbMock)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
		AvatarID: 1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "avatar_id", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, user.RoleID, user.AvatarID, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUser :one SELECT id, name, email, password, role_id, avatar_id, created_at, updated_at FROM users WHERE id = \$1`).
		WithArgs(user.ID).
		WillReturnRows(rows)

	u, err := repo.GetUser(user.ID)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.Password, u.Password)
	assert.Equal(t, user.RoleID, u.RoleID)
	assert.Equal(t, user.AvatarID, u.AvatarID)
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewUserRepository(config.DbMock)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
		AvatarID: 1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "avatar_id", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, user.RoleID, user.AvatarID, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUserByEmail :one SELECT id, name, email, password, role_id, avatar_id, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(user.Email).
		WillReturnRows(rows)

	u, err := repo.GetUserByEmail(user.Email)

	assert.NoError(t, err)

	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.Password, u.Password)
	assert.Equal(t, user.RoleID, u.RoleID)
	assert.Equal(t, user.AvatarID, u.AvatarID)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewUserRepository(config.DbMock)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
	}
	deleteUserSQL := `-- name: DeleteUser :exec DELETE FROM users WHERE id = \$1`
	mock.ExpectBegin()

	mock.ExpectExec(deleteUserSQL).
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = repo.Begin()
	assert.NoError(t, err)
	deletedUser := repo.DeleteUser(user.ID)
	assert.NoError(t, deletedUser)
	err = repo.Commit()
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock()
	repo := repository.NewUserRepository(config.DbMock)

	user := &entity.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "123456",
		RoleID:   1,
		AvatarID: 1,
	}

	updateUserSQL := `-- name: UpdateUser :exec UPDATE users SET name = \$1, email = \$2, password = \$3, role_id = \$4, avatar_id = \$5 WHERE id = \$6`
	mock.ExpectBegin()
	mock.ExpectExec(updateUserSQL).
		WithArgs(user.Name, user.Email, user.Password, user.RoleID, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = repo.Begin()
	assert.NoError(t, err)
	updateduser := repo.UpdateUser(user, 1)
	assert.NoError(t, updateduser)
	err = repo.Commit()
	assert.NoError(t, err)

}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock()
	repo := repository.NewUserRepository(config.DbMock)

	user := &entity.User{
		ID:       1,
		Name:     "Test",
		Email:    "test@teste.com",
		Password: "123456",
		RoleID:   1,
		AvatarID: 1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "avatar_id", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, user.RoleID, user.AvatarID, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetUsers :many SELECT id, name, email, password, role_id, avatar_id, created_at, updated_at FROM users ORDER BY id ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(int32(10), int32(0)).
		WillReturnRows(rows)

	users, err := repo.GetUsers(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(users))
	assert.Equal(t, user.Name, users[0].Name)
	assert.Equal(t, user.Email, users[0].Email)
	assert.Equal(t, user.Password, users[0].Password)
	assert.Equal(t, user.RoleID, users[0].RoleID)
	assert.Equal(t, user.AvatarID, users[0].AvatarID)
}
