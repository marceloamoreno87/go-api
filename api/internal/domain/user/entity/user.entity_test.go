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

	user.Name = ""
	err = user.Validate()
	assert.EqualError(t, err, "Name is required")

	user.Name = "Test User"
	user.Email = ""
	err = user.Validate()
	assert.EqualError(t, err, "Email is required")

	user.Email = "test@example.com"
	user.Password = ""
	err = user.Validate()
	assert.EqualError(t, err, "Password is required")

	user.Password = "password"
	user.RoleID = 0
	err = user.Validate()
	assert.EqualError(t, err, "Role is required")

	user.RoleID = 1
	err = user.Validate()
	assert.NoError(t, err)

	user.Name = ""
	user.Email = ""
	user.Password = ""
	user.RoleID = 0
	err = user.Validate()
	assert.EqualError(t, err, "Name is required")

}

func TestIsEmailValid(t *testing.T) {
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		RoleID:   1,
	}
	valid, err := user.IsEmailValid()
	assert.NoError(t, err)
	assert.True(t, valid)

	user.Email = "invalid email"
	valid, err = user.IsEmailValid()
	assert.EqualError(t, err, "Email is invalid")
	assert.False(t, valid)
}

func TestComparePassword(t *testing.T) {
	user, _ := entity.NewUser("Test User", "test@example.com", "password", 1)
	assert.True(t, user.ComparePassword("password"))
	assert.False(t, user.ComparePassword("wrongpassword"))
}
