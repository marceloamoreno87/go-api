package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreatePermission(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewPermissionRepository(db)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(permission.ID, permission.Name, permission.InternalName, permission.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: CreatePermission :one INSERT INTO permissions \( name, internal_name, description \) VALUES \( \$1, \$2, \$3 \) RETURNING id, name, internal_name, description, created_at, updated_at`).
		WithArgs(permission.Name, permission.InternalName, permission.Description).
		WillReturnRows(rows)

	err = rr.CreatePermission(permission)

	assert.NoError(t, err)

}

func TestGetPermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewPermissionRepository(db)

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

	r, err := rr.GetPermission(permission.ID)

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

	rr := repository.NewPermissionRepository(db)

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

	r, err := rr.GetPermissionByInternalName(permission.InternalName)

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

	rr := repository.NewPermissionRepository(db)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(permission.ID, permission.Name, permission.InternalName, permission.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: DeletePermission :one DELETE FROM permissions WHERE id = \$1 RETURNING id, name, internal_name, description, created_at, updated_at`).
		WithArgs(permission.ID).
		WillReturnRows(rows)

	err = rr.DeletePermission(permission.ID)

	assert.NoError(t, err)

}

func TestUpdatePermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewPermissionRepository(db)

	permission := &entity.Permission{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(permission.ID, permission.Name, permission.InternalName, permission.Description, time.Now(), time.Now())

	updatePermissionSQL := `-- name: UpdatePermission :one UPDATE permissions SET name = \$1, internal_name = \$2, description = \$3 WHERE id = \$4 RETURNING id, name, internal_name, description, created_at, updated_at`
	mock.ExpectQuery(updatePermissionSQL).
		WithArgs(permission.Name, permission.InternalName, permission.Description, permission.ID).
		WillReturnRows(rows)

	err = rr.UpdatePermission(permission, permission.ID)

	assert.NoError(t, err)

}

func TestGetPermissions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewPermissionRepository(db)

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

	permissions, err := rr.GetPermissions(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(permissions))
	assert.Equal(t, permission.Name, permissions[0].Name)
	assert.Equal(t, permission.InternalName, permissions[0].InternalName)
	assert.Equal(t, permission.Description, permissions[0].Description)

}
