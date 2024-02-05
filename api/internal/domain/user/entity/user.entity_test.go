package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Test User", "test@example.com", "password", 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.GetName())
	assert.Equal(t, "test@example.com", user.GetEmail())
	assert.True(t, user.ComparePassword("password"))
	assert.Equal(t, int32(1), user.GetRoleId())
}

func TestValidate(t *testing.T) {
	user := &User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		RoleId:   1,
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
	user.RoleId = 0
	err = user.Validate()
	assert.EqualError(t, err, "Role is required")

	user.RoleId = 1
	err = user.Validate()
	assert.NoError(t, err)

	user.Name = ""
	user.Email = ""
	user.Password = ""
	user.RoleId = 0
	err = user.Validate()
	assert.EqualError(t, err, "Name is required")

}

func TestIsEmailValid(t *testing.T) {
	user := &User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
		RoleId:   1,
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
	user, _ := NewUser("Test User", "test@example.com", "password", 1)
	assert.True(t, user.ComparePassword("password"))
	assert.False(t, user.ComparePassword("wrongpassword"))
}
