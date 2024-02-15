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

func TestNewInvalid(t *testing.T) {
	role, err := entity.NewRole("", "", "")

	assert.NotNil(t, err)
	assert.Nil(t, role)
}

func TestValidate(t *testing.T) {
	role := &entity.Role{Name: "name", InternalName: "test_role", Description: "description"}
	err := role.Validate()
	assert.Nil(t, err)
}

func TestGetID(t *testing.T) {
	role := &entity.Role{ID: 1}
	assert.Equal(t, int32(1), role.GetID())
}

func TestGetName(t *testing.T) {
	role := &entity.Role{Name: "name"}
	assert.Equal(t, "name", role.GetName())
}

func TestGetInternalName(t *testing.T) {
	role := &entity.Role{InternalName: "test_role"}
	assert.Equal(t, "test_role", role.GetInternalName())
}

func TestGetDescription(t *testing.T) {
	role := &entity.Role{Description: "description"}
	assert.Equal(t, "description", role.GetDescription())
}

func TestGetCreatedAt(t *testing.T) {
	role := &entity.Role{}
	assert.NotNil(t, role.GetCreatedAt())
}

func TestGetUpdatedAt(t *testing.T) {
	role := &entity.Role{}
	assert.NotNil(t, role.GetUpdatedAt())
}

func TestSetID(t *testing.T) {
	role := &entity.Role{}
	role.SetID(1)
	assert.Equal(t, int32(1), role.ID)
}

func TestSetName(t *testing.T) {
	role := &entity.Role{}
	role.SetName("name")
	assert.Equal(t, "name", role.Name)
}

func TestSetInternalName(t *testing.T) {
	role := &entity.Role{}
	role.SetInternalName("test_role")
	assert.Equal(t, "test_role", role.InternalName)
}

func TestSetDescription(t *testing.T) {
	role := &entity.Role{}
	role.SetDescription("description")
	assert.Equal(t, "description", role.Description)
}

func TestSetCreatedAt(t *testing.T) {
	role := &entity.Role{}
	assertionTime := time.Now()
	role.SetCreatedAt(assertionTime)
	assert.Equal(t, assertionTime, role.CreatedAt)
}

func TestSetUpdatedAt(t *testing.T) {
	role := &entity.Role{}
	assertionTime := time.Now()
	role.SetUpdatedAt(assertionTime)
	assert.Equal(t, assertionTime, role.UpdatedAt)
}
