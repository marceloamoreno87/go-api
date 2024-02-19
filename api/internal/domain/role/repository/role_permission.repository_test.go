package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateRolePermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewRolePermissionRepository(db)

	rolePermission := &entity.RolePermission{
		RoleID:        1,
		PermissionIDs: []int32{1},
	}

	createRolePermissionSQL := `-- name: CreateRolePermission :exec INSERT INTO role_permissions \( role_id, permission_id \) VALUES \( \$1, \$2 \)`
	mock.ExpectExec(createRolePermissionSQL).
		WithArgs(rolePermission.RoleID, rolePermission.PermissionIDs[0]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = rr.CreateRolePermission(rolePermission)

	assert.NoError(t, err)
}

func TestGetRolePermissions(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewRolePermissionRepository(db)

	rolePermission := &entity.RolePermission{
		ID:            1,
		RoleID:        1,
		PermissionIDs: []int32{1},
	}

	rows := sqlmock.NewRows([]string{"id", "role_id", "permission_id", "id", "name", "internal_name", "description", "created_at", "updated_at", "id_2", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(rolePermission.ID, rolePermission.RoleID, rolePermission.PermissionIDs[0], 1, "Test", "test", "test", time.Now(), time.Now(), 1, "Test_2", "test_2", "test_2", time.Now(), time.Now()).
		AddRow(rolePermission.ID, rolePermission.RoleID, rolePermission.PermissionIDs[0], 1, "Test", "test", "test", time.Now(), time.Now(), 1, "Test_2", "test_2", "test_2", time.Now(), time.Now())

	gerRolePermissionSQL := `-- name: GetRolePermissions :many SELECT role_permissions.id, role_id, permission_id, permissions.id, permissions.name, permissions.internal_name, permissions.description, permissions.created_at, permissions.updated_at, roles.id, roles.name, roles.internal_name, roles.description, roles.created_at, roles.updated_at FROM role_permissions INNER JOIN permissions ON role_permissions.permission_id = permissions.id INNER JOIN roles ON role_permissions.role_id = roles.id WHERE role_id = \$1 ORDER BY permission_id ASC`
	mock.ExpectQuery(gerRolePermissionSQL).
		WithArgs(rolePermission.RoleID).
		WillReturnRows(rows)

	_, err = rr.GetRolePermissions(rolePermission.RoleID)

	assert.NoError(t, err)
}

func TestUpdateRolePermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewRolePermissionRepository(db)

	rolePermission := &entity.RolePermission{
		RoleID:        1,
		PermissionIDs: []int32{1},
	}

	deleteRolePermissionSQL := `-- name: DeleteRolePermission :exec DELETE FROM role_permissions WHERE role_id = \$1`
	mock.ExpectExec(deleteRolePermissionSQL).
		WithArgs(rolePermission.RoleID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = rr.UpdateRolePermission(rolePermission, rolePermission.RoleID)

	createRolePermissionSQL := `-- name: CreateRolePermission :exec INSERT INTO role_permissions \( role_id, permission_id \) VALUES \( \$1, \$2 \)`
	mock.ExpectExec(createRolePermissionSQL).
		WithArgs(rolePermission.RoleID, rolePermission.PermissionIDs[0]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = rr.CreateRolePermission(rolePermission)

	assert.NoError(t, err)
}
