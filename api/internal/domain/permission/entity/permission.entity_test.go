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

func TestNewPermissionError(t *testing.T) {
	_, err := entity.NewPermission("", "", "")
	assert.NotNil(t, err)
}

func TestValidate(t *testing.T) {
	permission := &entity.Permission{Name: "name", InternalName: "test_permission", Description: "description"}
	err := permission.Validate()
	assert.Nil(t, err)
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
