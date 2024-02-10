package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := NewRoleRepository(db)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(role.ID, role.Name, role.InternalName, role.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: CreateRole :one INSERT INTO roles \( name, internal_name, description \) VALUES \( \$1, \$2, \$3 \) RETURNING id, name, internal_name, description, created_at, updated_at`).
		WithArgs(role.Name, role.InternalName, role.Description).
		WillReturnRows(rows)

	r, err := rr.CreateRole(role)

	assert.NoError(t, err)

	assert.Equal(t, role.Name, r.Name)
	assert.Equal(t, role.InternalName, r.InternalName)
	assert.Equal(t, role.Description, r.Description)
}

func TestGetRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := NewRoleRepository(db)

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

	r, err := rr.GetRole(role.ID)

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

	rr := NewRoleRepository(db)

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

	r, err := rr.GetRoleByInternalName(role.InternalName)

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

	rr := NewRoleRepository(db)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(role.ID, role.Name, role.InternalName, role.Description, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: DeleteRole :one DELETE FROM roles WHERE id = \$1 RETURNING id, name, internal_name, description, created_at, updated_at`).
		WithArgs(role.ID).
		WillReturnRows(rows)

	u, err := rr.DeleteRole(role.ID)

	assert.NoError(t, err)

	assert.Equal(t, role.Name, u.Name)
	assert.Equal(t, role.InternalName, u.InternalName)
	assert.Equal(t, role.Description, u.Description)

}

func TestUpdateRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := NewRoleRepository(db)

	role := &entity.Role{
		ID:           1,
		Name:         "Test",
		InternalName: "test_test",
		Description:  "testing",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "internal_name", "description", "created_at", "updated_at"}).
		AddRow(role.ID, role.Name, role.InternalName, role.Description, time.Now(), time.Now())

	updateRoleSQL := `-- name: UpdateRole :one UPDATE roles SET name = \$1, internal_name = \$2, description = \$3 WHERE id = \$4 RETURNING id, name, internal_name, description, created_at, updated_at`
	mock.ExpectQuery(updateRoleSQL).
		WithArgs(role.Name, role.InternalName, role.Description, 1).
		WillReturnRows(rows)

	r, err := rr.UpdateRole(role)

	assert.NoError(t, err)

	assert.Equal(t, role.Name, r.Name)
	assert.Equal(t, role.InternalName, r.InternalName)
	assert.Equal(t, role.Description, r.Description)

}

func TestGetRoles(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := NewRoleRepository(db)

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

	roles, err := rr.GetRoles(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(roles))
	assert.Equal(t, role.Name, roles[0].Name)
	assert.Equal(t, role.InternalName, roles[0].InternalName)
	assert.Equal(t, role.Description, roles[0].Description)

}
