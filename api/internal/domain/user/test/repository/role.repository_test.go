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

func TestCreateRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	config.NewDatabaseMock(db)
	repo := repository.NewRoleRepository(config.DbMock)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	createRoleSQL := `-- name: CreateRole :exec INSERT INTO roles \( name, internal_name, description \) VALUES \( \$1, \$2, \$3 \)`
	mock.ExpectBegin()
	mock.ExpectExec(createRoleSQL).
		WithArgs(role.Name, role.InternalName, role.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Begin()
	assert.NoError(t, err)
	newRole := repo.CreateRole(role)
	assert.NoError(t, newRole)
	err = repo.Commit()
	assert.NoError(t, err)
}

func TestGetRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewRoleRepository(config.DbMock)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(role.ID, role.Name, role.InternalName, role.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetRole :one SELECT id, name, internal_name, description, created_at, updated_at FROM roles WHERE id = \$1`).
		WithArgs(role.ID).
		WillReturnRows(rows)

	r, err := repo.GetRole(role.ID)

	assert.NoError(t, err)

	assert.Equal(t, role.Name, r.Name)
	assert.Equal(t, role.InternalName, r.InternalName)
	assert.Equal(t, role.Description, r.Description)

}

func TestGetRoleByInternalName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewRoleRepository(config.DbMock)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(role.ID, role.Name, role.InternalName, role.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetRoleByInternalName :one SELECT id, name, internal_name, description, created_at, updated_at FROM roles WHERE internal_name = \$1`).
		WithArgs("test_test").
		WillReturnRows(rows)

	r, err := repo.GetRoleByInternalName(role.InternalName)

	assert.NoError(t, err)

	assert.Equal(t, role.Name, r.Name)
	assert.Equal(t, role.InternalName, r.InternalName)
	assert.Equal(t, role.Description, r.Description)

}

func TestDeleteRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewRoleRepository(config.DbMock)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}
	deleteRoleSQL := `-- name: DeleteRole :exec DELETE FROM roles WHERE id = \$1`
	mock.ExpectBegin()

	mock.ExpectExec(deleteRoleSQL).
		WithArgs(role.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	err = repo.Begin()
	assert.NoError(t, err)
	deletedRole := repo.DeleteRole(role.ID)
	assert.NoError(t, deletedRole)
	err = repo.Commit()
	assert.NoError(t, err)
}

func TestUpdateRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewRoleRepository(config.DbMock)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	updateRoleSQL := `-- name: UpdateRole :exec UPDATE roles SET name = \$1, internal_name = \$2, description = \$3 WHERE id = \$4`
	mock.ExpectBegin()
	mock.ExpectExec(updateRoleSQL).
		WithArgs(role.Name, role.InternalName, role.Description, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	err = repo.Begin()
	assert.NoError(t, err)
	updatedRole := repo.UpdateRole(role, role.ID)
	assert.NoError(t, updatedRole)
	err = repo.Commit()
	assert.NoError(t, err)

}

func TestGetRoles(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewRoleRepository(config.DbMock)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(role.ID, role.Name, role.InternalName, role.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetRoles :many SELECT id, name, internal_name, description, created_at, updated_at FROM roles ORDER BY id ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(int32(10), int32(0)).
		WillReturnRows(rows)

	roles, err := repo.GetRoles(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(roles))
	assert.Equal(t, role.Name, roles[0].Name)
	assert.Equal(t, role.InternalName, roles[0].InternalName)
	assert.Equal(t, role.Description, roles[0].Description)

}
