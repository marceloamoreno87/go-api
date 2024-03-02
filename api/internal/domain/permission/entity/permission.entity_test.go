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

func TestNewPermissionEmpty(t *testing.T) {
	permission, err := entity.NewPermission("", "", "")
	assert.NotNil(t, err)
	assert.Nil(t, permission)
	assert.Equal(t, "[permission.entity.name]: Name is required, [permission.entity.internal_name]: Internal name is required, [permission.entity.description]: Description is required", err.Error())
}

func TestValidateName(t *testing.T) {
	permission, err := entity.NewPermission("", "test_permission", "description")
	assert.NotNil(t, err)
	assert.Nil(t, permission)
	assert.Equal(t, "[permission.entity.name]: Name is required", err.Error())
}

func TestValidateInternalName(t *testing.T) {
	permission, err := entity.NewPermission("name", "", "description")
	assert.NotNil(t, err)
	assert.Nil(t, permission)
	assert.Equal(t, "[permission.entity.internal_name]: Internal name is required", err.Error())
}

func TestValidateDescription(t *testing.T) {
	permission, err := entity.NewPermission("name", "test_permission", "")
	assert.NotNil(t, err)
	assert.Nil(t, permission)
	assert.Equal(t, "[permission.entity.description]: Description is required", err.Error())
}

func TestGetID(t *testing.T) {
	permission := &entity.Permission{ID: 1}
	assert.Equal(t, int32(1), permission.GetID())
}

func TestGetName(t *testing.T) {
	permission := &entity.Permission{Name: "name"}
	assert.Equal(t, "name", permission.GetName())
}

func TestGetInternalName(t *testing.T) {
	permission := &entity.Permission{InternalName: "test_permission"}
	assert.Equal(t, "test_permission", permission.GetInternalName())
}

func TestGetDescription(t *testing.T) {
	permission := &entity.Permission{Description: "description"}
	assert.Equal(t, "description", permission.GetDescription())
}

func TestGetCreatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assert.NotNil(t, permission.GetCreatedAt())
}

func TestGetUpdatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assert.NotNil(t, permission.GetUpdatedAt())
}

func TestSetID(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetID(1)
	assert.Equal(t, int32(1), permission.ID)
}

func TestSetName(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetName("name")
	assert.Equal(t, "name", permission.Name)
}

func TestSetInternalName(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetInternalName("test_permission")
	assert.Equal(t, "test_permission", permission.InternalName)
}

func TestSetDescription(t *testing.T) {
	permission := &entity.Permission{}
	permission.SetDescription("description")
	assert.Equal(t, "description", permission.Description)
}

func TestSetCreatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assertionTime := time.Now()
	permission.SetCreatedAt(assertionTime)
	assert.Equal(t, assertionTime, permission.CreatedAt)
}

func TestSetUpdatedAt(t *testing.T) {
	permission := &entity.Permission{}
	assertionTime := time.Now()
	permission.SetUpdatedAt(assertionTime)
	assert.Equal(t, assertionTime, permission.UpdatedAt)
}
