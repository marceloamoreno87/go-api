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

func TestCreatePermission(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	config.NewDatabaseMock(db)
	repo := repository.NewPermissionRepository(config.DbMock)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}
	createPermissionSQL := `-- name: CreatePermission :exec INSERT INTO permissions \( name, internal_name, description \) VALUES \( \$1, \$2, \$3 \)`
	mock.ExpectBegin()
	mock.ExpectExec(createPermissionSQL).
		WithArgs(permission.Name, permission.InternalName, permission.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Begin()
	assert.NoError(t, err)
	newPermission := repo.CreatePermission(permission)
	assert.NoError(t, newPermission)
	err = repo.Commit()
	assert.NoError(t, err)

}

func TestGetPermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewPermissionRepository(config.DbMock)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(permission.ID, permission.Name, permission.InternalName, permission.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetPermission :one SELECT id, name, internal_name, description, created_at, updated_at FROM permissions WHERE id = \$1 LIMIT 1`).
		WithArgs(permission.ID).
		WillReturnRows(rows)

	r, err := repo.GetPermission(permission.ID)

	assert.NoError(t, err)

	assert.Equal(t, permission.Name, r.Name)
	assert.Equal(t, permission.InternalName, r.InternalName)
	assert.Equal(t, permission.Description, r.Description)

}

func TestGetPermissionByInternalName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewPermissionRepository(config.DbMock)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(permission.ID, permission.Name, permission.InternalName, permission.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetPermissionByInternalName :one SELECT id, name, internal_name, description, created_at, updated_at FROM permissions WHERE internal_name = \$1`).
		WithArgs("test_test").
		WillReturnRows(rows)

	r, err := repo.GetPermissionByInternalName(permission.InternalName)

	assert.NoError(t, err)

	assert.Equal(t, permission.Name, r.Name)
	assert.Equal(t, permission.InternalName, r.InternalName)
	assert.Equal(t, permission.Description, r.Description)

}

func TestDeletePermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewPermissionRepository(config.DbMock)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	deletePermissionSQL := `-- name: DeletePermission :exec DELETE FROM permissions WHERE id = \$1`
	mock.ExpectBegin()
	mock.ExpectExec(deletePermissionSQL).
		WithArgs(permission.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Begin()
	assert.NoError(t, err)
	deletedPermission := repo.DeletePermission(permission.ID)
	assert.NoError(t, deletedPermission)
	err = repo.Commit()
	assert.NoError(t, err)

}

func TestUpdatePermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewPermissionRepository(config.DbMock)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	updatePermissionSQL := `-- name: UpdatePermission :exec UPDATE permissions SET name = \$1, internal_name = \$2, description = \$3 WHERE id = \$4`
	mock.ExpectBegin()
	mock.ExpectExec(updatePermissionSQL).
		WithArgs(permission.Name, permission.InternalName, permission.Description, permission.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Begin()
	assert.NoError(t, err)
	updatedPermission := repo.UpdatePermission(permission, permission.ID)
	assert.NoError(t, updatedPermission)
	err = repo.Commit()
	assert.NoError(t, err)

}

func TestGetPermissions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.NewDatabaseMock(db)
	repo := repository.NewPermissionRepository(config.DbMock)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(permission.ID, permission.Name, permission.InternalName, permission.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetPermissions :many SELECT id, name, internal_name, description, created_at, updated_at FROM permissions ORDER BY id ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(int32(10), int32(0)).
		WillReturnRows(rows)

	permissions, err := repo.GetPermissions(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(permissions))
	assert.Equal(t, permission.Name, permissions[0].Name)
	assert.Equal(t, permission.InternalName, permissions[0].InternalName)
	assert.Equal(t, permission.Description, permissions[0].Description)

}
