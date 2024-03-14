package entity_test

import (
	"testing"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("Test User", "test@example.com", "password", 1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.GetName())
	assert.Equal(t, "test@example.com", user.GetEmail())
	assert.True(t, user.ComparePassword("password"))
	assert.Equal(t, int32(1), user.GetRoleID())
}

func TestNewUserEmpty(t *testing.T) {
	user, err := entity.NewUser("", "", "", 0, 0)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.name]: Name is required, [user.entity.email]: Email is required, [user.entity.email]: Email is invalid, [user.entity.password]: Password is required, [user.entity.role_id]: Role is required, [user.entity.avatar_id]: Avatar is required", err.Error())
}

func TestValidateName(t *testing.T) {
	user, err := entity.NewUser("", "test@example.com", "password", 1, 1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.name]: Name is required", err.Error())
}

func TestValidateEmail(t *testing.T) {
	user, err := entity.NewUser("Test User", "", "password", 1, 1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.email]: Email is required, [user.entity.email]: Email is invalid", err.Error())
}

func TestValidatePassword(t *testing.T) {
	user, err := entity.NewUser("Test User", "test@example.com", "", 1, 1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.password]: Password is required", err.Error())
}

func TestValidateRoleID(t *testing.T) {
	user, err := entity.NewUser("Test User", "test@example.com", "password", 0, 1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.role_id]: Role is required", err.Error())
}

func TestValidateAvatarID(t *testing.T) {
	user, err := entity.NewUser("Test User", "test@example.com", "password", 1, 0)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.avatar_id]: Avatar is required", err.Error())
}

func TestValidateInvalidEmail(t *testing.T) {
	user, err := entity.NewUser("Test User", "test", "password", 1, 1)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "[user.entity.email]: Email is invalid", err.Error())
}

func TestComparePassword(t *testing.T) {
	user, _ := entity.NewUser("Test User", "test@example.com", "password", 1, 1)
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
