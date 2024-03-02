package entity_test

import (
	"testing"
	"time"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RolePermissionEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewRolePermission(t *testing.T) {
	rolePermission, err := RolePermissionEntity.NewRolePermission(1, []int32{1, 2, 3, 4})
	assert.NoError(t, err)
	assert.NotNil(t, rolePermission)
}

func TestNewRolePermissionError(t *testing.T) {
	_, err := RolePermissionEntity.NewRolePermission(1, []int32{})
	assert.Error(t, err)
}

func TestRolePermissionValidateEmpty(t *testing.T) {
	rolePermission, err := RolePermissionEntity.NewRolePermission(0, []int32{})
	assert.Error(t, err)
	assert.Nil(t, rolePermission)
	assert.Equal(t, "[role_permission.entity.role_id]: RoleID is required, [role_permission.entity.permission_ids]: PermissionIDs is required", err.Error())
}

func TestRolePermissionValidate(t *testing.T) {
	rolePermission, err := RolePermissionEntity.NewRolePermission(1, []int32{1, 2, 3, 4})
	assert.NoError(t, err)
	assert.NotNil(t, rolePermission)
}

func TestRolePermissionValidateRoleID(t *testing.T) {
	rolePermission, err := RolePermissionEntity.NewRolePermission(0, []int32{1, 2, 3, 4})
	assert.Error(t, err)
	assert.Nil(t, rolePermission)
	assert.Equal(t, "[role_permission.entity.role_id]: RoleID is required", err.Error())
}

func TestRolePermissionValidatePermissionIDs(t *testing.T) {
	rolePermission, err := RolePermissionEntity.NewRolePermission(1, []int32{})
	assert.Error(t, err)
	assert.Nil(t, rolePermission)
	assert.Equal(t, "[role_permission.entity.permission_ids]: PermissionIDs is required", err.Error())
}

func TestGetRoleID(t *testing.T) {
	rolePermission := &RolePermissionEntity.RolePermission{RoleID: 1}
	assert.Equal(t, int32(1), rolePermission.GetRoleID())
}

func TestGetPermissionIDs(t *testing.T) {
	rolePermission := &RolePermissionEntity.RolePermission{PermissionIDs: []int32{1, 2, 3, 4}}
	assert.Equal(t, []int32{1, 2, 3, 4}, rolePermission.GetPermissionIDs())
}

func TestSetRole(t *testing.T) {
	role := &RoleEntity.Role{
		ID:           1,
		Name:         "name",
		InternalName: "test_role",
		Description:  "description",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	rolePermission := &RolePermissionEntity.RolePermission{
		RoleID:        1,
		PermissionIDs: []int32{1, 2, 3, 4},
	}
	rolePermission.SetRole(role)
	assert.Equal(t, int32(1), rolePermission.RoleID)
}

func TestGetRole(t *testing.T) {
	role := &RoleEntity.Role{
		ID:           1,
		Name:         "name",
		InternalName: "test_role",
		Description:  "description",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	rolePermission := &RolePermissionEntity.RolePermission{
		RoleID:        1,
		PermissionIDs: []int32{1, 2, 3, 4},
	}
	rolePermission.SetRole(role)
	assert.Equal(t, role, rolePermission.GetRole())
}

func TestSetPermission(t *testing.T) {
	permissions := []*PermissionEntity.Permission{
		{
			ID:          1,
			Name:        "name",
			Description: "description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	rolePermission := &RolePermissionEntity.RolePermission{
		RoleID:        1,
		PermissionIDs: []int32{1, 2, 3, 4},
	}
	rolePermission.SetPermission(permissions)
	assert.Equal(t, permissions, rolePermission.GetPermission())
}

func TestGetPermission(t *testing.T) {
	permissions := []*PermissionEntity.Permission{
		{
			ID:          1,
			Name:        "name",
			Description: "description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	rolePermission := &entity.RolePermission{
		RoleID:        1,
		PermissionIDs: []int32{1, 2, 3, 4},
	}
	rolePermission.SetPermission(permissions)
	assert.Equal(t, permissions, rolePermission.GetPermission())
}
