package entity_test

import (
	"testing"
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewRole(t *testing.T) {
	role, err := entity.NewRole("name", "test_role", "description")
	assert.Nil(t, err)
	assert.NotNil(t, role)

	assert.Equal(t, "name", role.Name)
	assert.Equal(t, "test_role", role.InternalName)
	assert.Equal(t, "description", role.Description)
}

func TestNewRole_Invalid(t *testing.T) {
	role, err := entity.NewRole("", "", "")

	assert.NotNil(t, err)
	assert.Nil(t, role)
}

func TestNewRole_InvalidName(t *testing.T) {
	role, err := entity.NewRole("", "test_role", "description")

	assert.NotNil(t, err)
	assert.Nil(t, role)
}

func TestNewRole_InvalidInternalName(t *testing.T) {
	role, err := entity.NewRole("name", "", "description")

	assert.NotNil(t, err)
	assert.Nil(t, role)
}

func TestNewRole_InvalidDescription(t *testing.T) {
	role, err := entity.NewRole("name", "test_role", "")

	assert.NotNil(t, err)
	assert.Nil(t, role)
}

func TestRole_GetID(t *testing.T) {
	role := &entity.Role{ID: 1}
	assert.Equal(t, int32(1), role.GetID())
}

func TestRole_GetName(t *testing.T) {
	role := &entity.Role{Name: "name"}
	assert.Equal(t, "name", role.GetName())
}

func TestRole_GetInternalName(t *testing.T) {
	role := &entity.Role{InternalName: "test_role"}
	assert.Equal(t, "test_role", role.GetInternalName())
}

func TestRole_GetDescription(t *testing.T) {
	role := &entity.Role{Description: "description"}
	assert.Equal(t, "description", role.GetDescription())
}

func TestRole_GetCreatedAt(t *testing.T) {
	role := &entity.Role{}
	assert.NotNil(t, role.GetCreatedAt())
}

func TestRole_GetUpdatedAt(t *testing.T) {
	role := &entity.Role{}
	assert.NotNil(t, role.GetUpdatedAt())
}

func TestRole_SetID(t *testing.T) {
	role := &entity.Role{}
	role.SetID(1)
	assert.Equal(t, int32(1), role.ID)
}

func TestRole_SetName(t *testing.T) {
	role := &entity.Role{}
	role.SetName("name")
	assert.Equal(t, "name", role.Name)
}

func TestRole_SetInternalName(t *testing.T) {
	role := &entity.Role{}
	role.SetInternalName("test_role")
	assert.Equal(t, "test_role", role.InternalName)
}

func TestRole_SetDescription(t *testing.T) {
	role := &entity.Role{}
	role.SetDescription("description")
	assert.Equal(t, "description", role.Description)
}

func TestRole_SetCreatedAt(t *testing.T) {
	role := &entity.Role{}
	assertionTime := time.Now()
	role.SetCreatedAt(assertionTime)
	assert.Equal(t, assertionTime, role.CreatedAt)
}

func TestRole_SetUpdatedAt(t *testing.T) {
	role := &entity.Role{}
	assertionTime := time.Now()
	role.SetUpdatedAt(assertionTime)
	assert.Equal(t, assertionTime, role.UpdatedAt)
}
