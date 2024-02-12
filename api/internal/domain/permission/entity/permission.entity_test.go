package entity_test

import (
	"testing"
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewPermission(t *testing.T) {
	permission, err := entity.NewPermission("name", "test_permission", "description")

	assert.Nil(t, err)
	assert.NotNil(t, permission)

	assert.Equal(t, "name", permission.Name)
	assert.Equal(t, "test_permission", permission.InternalName)
	assert.Equal(t, "description", permission.Description)

}

func TestNewPermission_Invalid(t *testing.T) {
	permission, err := entity.NewPermission("", "", "")

	assert.NotNil(t, err)
	assert.Nil(t, permission)
}

func TestNewPermission_InvalidName(t *testing.T) {
	permission, err := entity.NewPermission("", "test_permission", "description")

	assert.NotNil(t, err)
	assert.Nil(t, permission)
}

func TestNewPermission_InvalidInternalName(t *testing.T) {
	permission, err := entity.NewPermission("name", "", "description")

	assert.NotNil(t, err)
	assert.Nil(t, permission)
}

func TestNewPermission_InvalidDescription(t *testing.T) {
	permission, err := entity.NewPermission("name", "test_permission", "")

	assert.NotNil(t, err)
	assert.Nil(t, permission)
}

func TestPermission_GetID(t *testing.T) {
	permission := &entity.Permission{ID: 1}
	assert.Equal(t, 1, permission.GetID())
}

func TestPermission_GetName(t *testing.T) {
	permission := &entity.Permission{Name: "name"}
	assert.Equal(t, "name", permission.GetName())
}

func TestPermission_GetInternalName(t *testing.T) {
	permission := &entity.Permission{InternalName: "test_permission"}
	assert.Equal(t, "test_permission", permission.GetInternalName())
}

func TestPermission_GetDescription(t *testing.T) {
	permission := &entity.Permission{Description: "description"}
	assert.Equal(t, "description", permission.GetDescription())
}

func TestPermission_GetCreatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assert.NotNil(t, permission.GetCreatedAt())
}

func TestPermission_GetUpdatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assert.NotNil(t, permission.GetUpdatedAt())
}

func TestPermission_SetID(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetID(1)
	assert.Equal(t, 1, permission.ID)
}

func TestPermission_SetName(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetName("name")
	assert.Equal(t, "name", permission.Name)
}

func TestPermission_SetInternalName(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetInternalName("test_permission")
	assert.Equal(t, "test_permission", permission.InternalName)
}

func TestPermission_SetDescription(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetDescription("description")
	assert.Equal(t, "description", permission.Description)
}

func TestPermission_SetCreatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assertionTime := time.Now()
	permission.SetCreatedAt(assertionTime)
	assert.Equal(t, assertionTime, permission.CreatedAt)
}

func TestPermission_SetUpdatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assertionTime := time.Now()
	permission.SetUpdatedAt(assertionTime)
	assert.Equal(t, assertionTime, permission.UpdatedAt)
}
