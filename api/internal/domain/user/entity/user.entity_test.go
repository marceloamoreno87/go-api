package entity_test

import (
	"testing"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("Test User", "test@example.com", "password", 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.GetName())
	assert.Equal(t, "test@example.com", user.GetEmail())
	assert.True(t, user.ComparePassword("password"))
	assert.Equal(t, int32(1), user.GetRoleID())
}

func TestValidate(t *testing.T) {
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		RoleID:   1,
	}
	err := user.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidEmail(t *testing.T) {
	user := &entity.User{
		Name:     "Test User",
		Email:    "test",
		Password: "password",
		RoleID:   1,
	}
	err := user.Validate()
	assert.Error(t, err)
}

func TestComparePassword(t *testing.T) {
	user, _ := entity.NewUser("Test User", "test@example.com", "password", 1)
	assert.True(t, user.ComparePassword("password"))
	assert.False(t, user.ComparePassword("wrongpassword"))
}

func TestGetName(t *testing.T) {
	user := &entity.User{Name: "Test User"}
	assert.Equal(t, "Test User", user.GetName())
}

func TestGetEmail(t *testing.T) {
	user := &entity.User{Email: "test@test.com"}
	assert.Equal(t, "test@test.com", user.GetEmail())
}

func TestGetPassword(t *testing.T) {
	user := &entity.User{Password: "password"}
	assert.Equal(t, "password", user.GetPassword())
}

func TestGetRoleID(t *testing.T) {
	user := &entity.User{RoleID: 1}
	assert.Equal(t, int32(1), user.GetRoleID())
}

func TestGetID(t *testing.T) {
	user := &entity.User{ID: 1}
	assert.Equal(t, int32(1), user.GetID())
}

func TestSetID(t *testing.T) {
	user := &entity.User{}
	user.SetID(1)
	assert.Equal(t, int32(1), user.GetID())
}

func TestSetRoleID(t *testing.T) {
	user := &entity.User{}
	user.SetRoleID(1)
	assert.Equal(t, int32(1), user.GetRoleID())
}

func TestSetPassword(t *testing.T) {
	user := &entity.User{}
	user.SetPassword("password")
	assert.NotEmpty(t, user.GetPassword())
}
